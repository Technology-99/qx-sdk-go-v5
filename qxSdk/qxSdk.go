package qxSdk

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCms"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxConfig"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxMas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxSas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTpas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

//go:embed "version"
var VersionF embed.FS

type QxSdk struct {
	Version string

	ctx        context.Context    // 控制退出
	cancel     context.CancelFunc // 取消函数
	wg         sync.WaitGroup     // 等待后台任务退出
	isShutdown bool               // 标记 SDK 是否已经关闭

	Cli *qxCli.QxClient

	// note: 消息服务
	MasService qxMas.MasService
	// note: 存储服务
	SasService qxSas.SasService
	// note: 定时任务队列服务
	CtasService qxCtas.CtasService
	// note: 第三方聚合服务
	TpasService qxTpas.TpasService
	// note: 内容管理服务
	CmsService qxCms.CmsService
}

func NewQxSdk(AccessKeyId, AccessKeySecret, Endpoint string) *QxSdk {

	c := qxConfig.DefaultConfig(AccessKeyId, AccessKeySecret, Endpoint)

	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	qxClient := qxCli.NewQxClient(ctx, c)

	sdk := &QxSdk{
		Version:     string(versionFile),
		Cli:         qxClient,
		ctx:         ctx,
		cancel:      cancel,
		MasService:  qxMas.NewMasService(qxClient),
		SasService:  qxSas.NewSasService(qxClient),
		CtasService: qxCtas.NewCtasService(qxClient),
		TpasService: qxTpas.NewTpasService(qxClient),
		CmsService:  qxCms.NewCmsService(qxClient),
	}
	sdk.AutoAuth()
	return sdk
}
func (s *QxSdk) GetVersion() string {
	return s.Version
}

func (s *QxSdk) AutoAuth() *QxSdk {
	s, _ = s.AuthHealthZ().AuthLogin()
	go s.AutoRefresh()
	return s
}

func (s *QxSdk) AutoRefresh() *QxSdk {
	if s.Cli.Config.AutoRefreshToken {
		s.wg.Add(1)

		go func() {
			defer s.wg.Done() // 确保任务退出时通知 WaitGroup
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					if s.Cli.RetryTimes > s.Cli.Config.MaxRetryTimes {
						// note: close auto refresh
						s.Cli.Config.AutoRefreshToken = false
						logx.Errorf("RefreshToken fail: %+v", qxTypes.ErrMaxErrTimes)
						break
					}
					if _, err := s.AuthRefresh(); err != nil {
						logx.Errorf("RefreshToken fail: %+v", err)
						s.AuthFail(err, "step1")
						time.Sleep(time.Second)
						continue
						//return errs
					}
				case <-s.ctx.Done():
					fmt.Println("AutoRefresh stopped.")
					return
				}
			}
		}()
	}
	return s
}

// note: sdk auth api
func (s *QxSdk) AuthHealthZ() *QxSdk {
	reqFn := s.Cli.EasyNewRequest(s.Cli.Context, "/healthz", "GET", nil)
	result, err := reqFn()
	if err != nil {
		logx.Errorf("healthz request error: %v", err)
		return nil
	}
	res := qxTypes.HealthzResp{}
	_ = json.Unmarshal(result, &res)
	if res.Code == response.SUCCESS {
		logx.Infof("sdk healthz success")
		s.Cli.Status = qxCli.STATUS_READY
	} else {
		logx.Errorf("healthz request error: %v", res.Msg)
		panic(res.Msg)
	}
	return s
}

func (s *QxSdk) AuthLogin() (*QxSdk, error) {
	logx.Infof("打印sdk的状态: %s", s.FormatQxSdkStatus())
	if s.Cli.Status == qxCli.STATUS_NOT_READY {
		logx.Errorf("sdk not ready")
		return s, qxTypes.ErrNotReady
	}

	reqFn := s.Cli.EasyNewRequest(s.Cli.Context, "/auth/sign", "POST", &qxTypes.ApiSignReq{
		AccessKey:    s.Cli.Config.AccessKeyId,
		AccessSecret: s.Cli.Config.AccessKeySecret,
	})
	result, err := reqFn()
	if err != nil {
		logx.Errorf("api sign error: %v", err)
		if s.Cli.Config.AutoRetry {
			if s.Cli.RetryTimes > s.Cli.Config.MaxRetryTimes {
				s.Cli.Status = qxCli.STATUS_NOT_READY
				logx.Errorf("sdk fail max times: %v", err)
				return nil, err
			} else {
				s.AuthFail(err, "step2")
				return s.AuthLogin()
			}
		} else {
			s.AuthFail(err, "step3")
		}
	}
	res := qxTypes.ApiSignResp{}
	_ = json.Unmarshal(result, &res)
	if res.Code == response.SUCCESS {
		logx.Infof("sdk api sign success")
		if s.Cli.Config.Debug {
			logx.Infof("sdk api sign result: %s", string(result))
		}
		// note: 记录access token
		s.Cli.AccessToken = res.Data.AccessToken
		s.Cli.AccessTokenExpires = res.Data.AccessExpiresAt
		s.Cli.RefreshToken = res.Data.RefreshToken
		s.Cli.RefreshTokenExpires = res.Data.RefreshExpiresAt
		s.AuthSuccess()
	} else {
		if s.Cli.Config.AutoRetry {
			if s.Cli.RetryTimes > s.Cli.Config.MaxRetryTimes {
				s.Cli.Status = qxCli.STATUS_NOT_READY
				logx.Errorf("sdk fail max times: %v", err)
				return nil, err
			} else {
				s.AuthFail(errors.New(res.Msg), "step3")
				return s.AuthLogin()
			}
		} else {
			s.AuthFail(errors.New(res.Msg), "step4")
		}
	}

	return s, nil
}

func (s *QxSdk) AuthRefresh() (*QxSdk, error) {
	if s.Cli.Config.Debug {
		logx.Infof("打印sdk的状态: %s", s.FormatQxSdkStatus())
	}
	// note: 如果链接未准备好，直接返回
	if s.Cli.Status == qxCli.STATUS_NOT_READY {
		return nil, qxTypes.ErrNotReady
	}

	nowTime := time.Now()

	// note: 判断accessToken过期了没
	if (s.Cli.AccessTokenExpires - s.Cli.Config.Deadline) >= nowTime.Unix() {
		if s.Cli.Config.Debug {
			logx.Infof("accessToken没过期，过期时间为: %s, 当前时间为: %s", time.Unix(s.Cli.AccessTokenExpires, 0).Format(time.DateTime), nowTime.Format(time.DateTime))
		}
		// note: 没过期，直接返回
		return s, nil
	}
	if (s.Cli.RefreshTokenExpires - s.Cli.Config.Deadline) >= nowTime.Unix() {
		if s.Cli.Config.Debug {
			logx.Infof("accessToken过期了，过期时间为: %s, 但是refreshToken没过期，过期时间为: %s, 当前时间为: %s", time.Unix(s.Cli.AccessTokenExpires, 0).Format(time.DateTime), time.Unix(s.Cli.RefreshTokenExpires, 0).Format(time.DateTime), nowTime.Format(time.DateTime))
		}
		// note: refreshToken没过期，但是accessToken过期了
		reqFn := s.Cli.EasyNewRequest(s.Cli.Context, "/auth/refresh", "POST", &qxTypes.ApiRefreshReq{
			AccessKey:    s.Cli.Config.AccessKeyId,
			RefreshToken: s.Cli.RefreshToken,
		})
		result, err := reqFn()
		if err != nil {
			logx.Errorf("api refresh error: %v", err)
			if s.Cli.Config.AutoRetry {
				if s.Cli.RetryTimes > s.Cli.Config.MaxRetryTimes {
					s.Cli.Status = qxCli.STATUS_NOT_READY
					logx.Errorf("sdk fail max times: %v", err)
					return nil, err
				} else {
					s.AuthFail(err, "step5")
					return s.AuthRefresh()
				}
			} else {
				s.AuthFail(err, "step6")
			}
		}
		res := qxTypes.ApiRefreshResp{}
		_ = json.Unmarshal(result, &res)
		if res.Code == response.SUCCESS {
			logx.Infof("api refresh success")
			if s.Cli.Config.Debug {
				logx.Infof("sdk api refresh result: %s", string(result))
			}
			// note: 记录access token
			s.Cli.AccessToken = res.Data.AccessToken
			s.Cli.AccessTokenExpires = res.Data.AccessExpiresAt
			s.AuthSuccess()
		} else {
			if s.Cli.Config.AutoRetry {
				if s.Cli.RetryTimes > s.Cli.Config.MaxRetryTimes {
					s.Cli.Status = qxCli.STATUS_NOT_READY
					logx.Errorf("sdk fail max times: %v", err)
					return nil, err
				} else {
					s.AuthFail(errors.New(res.Msg), "step7")
					return s.AuthRefresh()
				}
			} else {
				logx.Infof("api refresh fail: %s", res.Msg)
				s.AuthFail(errors.New(res.Msg), "step8")
			}
		}
	} else {
		// note: refreshToken过期了
		logx.Errorf("refreshToken 过期了")
		s.Cli.Status = qxCli.STATUS_NOT_READY
		return s.AuthHealthZ().AuthLogin()
	}
	return s, nil
}

func (s *QxSdk) AuthSuccess() {
	s.Cli.RetryTimes = 0
	s.Cli.Status = qxCli.STATUS_LOGINED
}

func (s *QxSdk) AuthFail(err error, step string) {
	if s.Cli.Config.AutoRetry {
		logx.Infof("sdk auth  step: %s, fail: %v", step, err)
		s.Cli.RetryTimes++
	} else {
		s.Cli.Status = qxCli.STATUS_NOT_READY
		panic(err)
	}
}

func (s *QxSdk) FormatQxSdkStatus() string {
	switch s.Cli.Status {
	case qxCli.STATUS_READY:
		return "已就绪"
	case qxCli.STATUS_LOGINED:
		return "已登录"
	case qxCli.STATUS_NOT_READY:
		return "未就绪"
	}
	return "未知状态"
}

func (s *QxSdk) Destroy() {
	if s.isShutdown {
		fmt.Println("SDK already shutdown.")
		return
	}

	fmt.Println("Shutting down SDK...")
	s.isShutdown = true

	// 通知后台任务退出
	s.cancel()

	// 等待所有 goroutine 退出
	s.wg.Wait()

	// 关闭 HTTP 客户端
	s.Cli.Client.CloseIdleConnections()

	fmt.Println("SDK resources released.")
}
