package sdk

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/cli"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/config"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/msg"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/req"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/resp"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/types"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

//go:embed "version"
var VersionF embed.FS

type Sdk struct {
	Version string

	ctx        context.Context    // 控制退出
	cancel     context.CancelFunc // 取消函数
	wg         sync.WaitGroup     // 等待后台任务退出
	isShutdown bool               // 标记 SDK 是否已经关闭

	Cli *cli.QxClient

	// note: 消息服务
	MsgService msg.MsgService
}

func NewSdk(AccessKeyId, AccessKeySecret, Endpoint string) *Sdk {

	c := config.DefaultConfig(AccessKeyId, AccessKeySecret, Endpoint)

	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	qxClient := cli.NewQxClient(ctx, c)

	sdk := &Sdk{
		Version:    string(versionFile),
		Cli:        qxClient,
		MsgService: msg.NewMsgService(qxClient),
		ctx:        ctx,
		cancel:     cancel,
	}
	sdk.AutoAuth()
	return sdk
}
func (s *Sdk) GetVersion() string {
	return s.Version
}

func (s *Sdk) AutoAuth() *Sdk {
	s, _ = s.AuthHealthZ().AuthLogin()
	go s.AutoRefresh()
	return s
}

func (s *Sdk) AutoRefresh() *Sdk {
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
						logx.Errorf("RefreshToken fail: %+v", types.ErrMaxErrTimes)
						break
					}
					if _, err := s.AuthRefresh(); err != nil {
						logx.Errorf("RefreshToken fail: %+v", err)
						s.AuthFail(err)
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
func (s *Sdk) AuthHealthZ() *Sdk {
	reqFn := s.Cli.EasyNewRequest(s.Cli.Context, "/healthz", "GET", nil)
	result, err := reqFn()
	if err != nil {
		logx.Errorf("healthz request error: %v", err)
		return nil
	}
	res := resp.DefaultResponseBase{}
	_ = json.Unmarshal(result, &res)
	if res.Code == response.SUCCESS {
		logx.Infof("sdk healthz success")
		s.Cli.Status = cli.STATUS_READY
	} else {
		panic(res.Msg)
	}
	return s
}

func (s *Sdk) AuthLogin() (*Sdk, error) {
	logx.Infof("打印sdk的状态: %s", s.FormatSdkStatus())
	if s.Cli.Status == cli.STATUS_NOT_READY {
		logx.Errorf("sdk not ready")
		return s, types.ErrNotReady
	}

	reqFn := s.Cli.EasyNewRequest(s.Cli.Context, "/auth/sign", "POST", &req.QxV5ApisApiSignReq{
		AccessKey:    s.Cli.Config.AccessKeyId,
		AccessSecret: s.Cli.Config.AccessKeySecret,
	})
	result, err := reqFn()
	if err != nil {
		logx.Errorf("api sign error: %v", err)
		if s.Cli.Config.AutoRetry {
			if s.Cli.RetryTimes > s.Cli.Config.MaxRetryTimes {
				s.Cli.Status = cli.STATUS_NOT_READY
				panic(types.ErrMaxErrTimes)
			} else {
				s.AuthFail(err)
				return s.AuthLogin()
			}
		} else {
			s.AuthFail(err)
		}
	}
	res := resp.QxV5ApisApiSignResp{}
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
				s.Cli.Status = cli.STATUS_NOT_READY
				panic(types.ErrMaxErrTimes)
			} else {
				s.AuthFail(errors.New(res.Msg))
				return s.AuthLogin()
			}
		} else {
			s.AuthFail(errors.New(res.Msg))
		}
	}

	return s, nil
}

func (s *Sdk) AuthRefresh() (*Sdk, error) {
	if s.Cli.Config.Debug {
		logx.Infof("打印sdk的状态: %s", s.FormatSdkStatus())
	}
	// note: 如果链接未准备好，直接返回
	if s.Cli.Status == cli.STATUS_NOT_READY {
		return nil, types.ErrNotReady
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
		reqFn := s.Cli.EasyNewRequest(s.Cli.Context, "/auth/refresh", "POST", &req.QxV5ApisApiRefreshReq{
			AccessKey:    s.Cli.Config.AccessKeyId,
			RefreshToken: s.Cli.RefreshToken,
		})
		result, err := reqFn()
		if err != nil {
			logx.Errorf("api refresh error: %v", err)
			if s.Cli.Config.AutoRetry {
				if s.Cli.RetryTimes > s.Cli.Config.MaxRetryTimes {
					s.Cli.Status = cli.STATUS_NOT_READY
					panic(types.ErrMaxErrTimes)
				} else {
					s.AuthFail(err)
					return s.AuthRefresh()
				}
			} else {
				s.AuthFail(err)
			}
		}
		res := resp.QxV5ApisApiRefreshResp{}
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
					s.Cli.Status = cli.STATUS_NOT_READY
					panic(types.ErrMaxErrTimes)
				} else {
					s.AuthFail(errors.New(res.Msg))
					return s.AuthRefresh()
				}
			} else {
				s.AuthFail(errors.New(res.Msg))
			}
		}
	} else {
		// note: refreshToken过期了
		logx.Errorf("refreshToken 过期了")
		s.Cli.Status = cli.STATUS_NOT_READY
		return s.AuthHealthZ().AuthLogin()
	}
	return s, nil
}

func (s *Sdk) AuthSuccess() {
	s.Cli.RetryTimes = 0
	s.Cli.Status = cli.STATUS_LOGINED
}

func (s *Sdk) AuthFail(err error) {
	if s.Cli.Config.AutoRetry {
		s.Cli.RetryTimes++
	} else {
		s.Cli.Status = cli.STATUS_NOT_READY
		panic(err)
	}
}

func (s *Sdk) FormatSdkStatus() string {
	switch s.Cli.Status {
	case cli.STATUS_READY:
		return "已就绪"
	case cli.STATUS_LOGINED:
		return "已登录"
	case cli.STATUS_NOT_READY:
		return "未就绪"
	}
	return "未知状态"
}

func (s *Sdk) Destroy() {
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
