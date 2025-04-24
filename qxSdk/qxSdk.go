package qxSdk

import (
	"context"
	"crypto/ecdh"
	"crypto/rand"
	"embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxBase"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCcs"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxConfig"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxErrs"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxKms"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxMas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxParser"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxSas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTpas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
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

	QxCtx *qxCtx.QxCtx
	// note: 标记sdk的状态
	Status int

	// note: 锁
	mu sync.Mutex

	// note: 琼霄自身的服务
	qxBase.QxBaseService

	// note: 云端配置存储服务
	CcsService qxCcs.CcsService
	// note: 消息服务
	MasService qxMas.MasService
	// note: 存储服务
	SasService qxSas.SasService
	// note: 定时任务队列服务
	CtasService qxCtas.CtasService
	// note: 第三方聚合服务
	TpasService qxTpas.TpasService
	// note: 密钥管理服务
	KmsService qxKms.KmsService
}

func NewQxSdk(c *qxConfig.Config) *QxSdk {
	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	qxClient := qxCli.NewQxClient(ctx, c)
	qxC := qxCtx.NewQxCtx(qxClient)

	sdk := &QxSdk{
		Version:       string(versionFile),
		QxCtx:         qxC,
		Status:        SdkStatusNotReady,
		ctx:           ctx,
		cancel:        cancel,
		QxBaseService: qxBase.NewQxBaseService(qxC),
		CcsService:    qxCcs.NewCcsService(qxC),
		KmsService:    qxKms.NewKmsService(qxC),
		MasService:    qxMas.NewMasService(qxC),
		SasService:    qxSas.NewSasService(qxC),
		CtasService:   qxCtas.NewCtasService(qxC),
		TpasService:   qxTpas.NewTpasService(qxC),
	}
	sdk.Init()
	return sdk
}

func NewDefaultQxSdk(AccessKeyId, AccessKeySecret, Endpoint string) *QxSdk {

	c := qxConfig.DefaultConfig(AccessKeyId, AccessKeySecret, Endpoint)

	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	qxClient := qxCli.NewQxClient(ctx, c)
	qxC := qxCtx.NewQxCtx(qxClient)

	sdk := &QxSdk{
		Version:     string(versionFile),
		QxCtx:       qxC,
		Status:      SdkStatusNotReady,
		ctx:         ctx,
		cancel:      cancel,
		CcsService:  qxCcs.NewCcsService(qxC),
		KmsService:  qxKms.NewKmsService(qxC),
		MasService:  qxMas.NewMasService(qxC),
		SasService:  qxSas.NewSasService(qxC),
		CtasService: qxCtas.NewCtasService(qxC),
		TpasService: qxTpas.NewTpasService(qxC),
	}

	sdk.Init()
	return sdk
}

func (s *QxSdk) Init() *QxSdk {
	s, _ = s.HealthZ().AuthLogin()
	s.KeyExChange()
	go s.AutoHealthZ()
	go s.AutoKeyExChange()
	go s.AutoRefresh()
	return s
}

// note: 验证远端的健康状态
func (s *QxSdk) HealthZ() *QxSdk {
	reqFn := s.EasyUnLoginRequest(s.QxCtx.Cli.Context, "/healthz", "GET", nil)
	result, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk healthz request error: %v", err)
		return nil
	}
	res := qxTypes.QxClientHealthzResp{}
	_ = json.Unmarshal(result, &res)
	if res.Code == qxCodes.QxEngineStatusOK {
		logx.Infof("qx sdk healthz success")
		if s.Status <= SdkStatusRemoteUnHealthy {
			s.Status = SdkStatusRemoteHealthy
		}
	} else {
		logx.Errorf("qx sdk healthz request error: %v", res.Msg)
		panic(res.Msg)
	}
	return s
}

func (s *QxSdk) AuthLogin() (*QxSdk, error) {
	logx.Infof("打印sdk的状态: %s", s.FormatQxSdkStatus())
	if s.Status < SdkStatusRemoteHealthy {
		logx.Errorf("sdk not ready")
		return s, qxErrs.ErrRemoteNotHealthy
	}
	reqFn := s.EasyUnLoginRequest(s.QxCtx.Cli.Context, "/auth/sign", "POST", &qxTypes.QxClientApiSignReq{
		AccessKey:    s.QxCtx.Cli.Config.AccessKeyId,
		AccessSecret: s.QxCtx.Cli.Config.AccessKeySecret,
	})
	result, err := reqFn()
	if err != nil {
		if s.QxCtx.Cli.Config.AutoRetry {
			if s.QxCtx.Cli.RetryTimes > s.QxCtx.Cli.Config.MaxRetryTimes {
				s.Status = SdkStatusSignMaxTimes
				logx.Errorf("qx sdk: fail max times: %v", err)
				return nil, err
			} else {
				logx.Errorf("qx sdk: sign failed, next try, err: %v, ", err)
				s.QxCtx.Cli.RetryTimes++
				return s.AuthLogin()
			}
		} else {
			s.Status = SdkStatusLoginFailed
			logx.Infof("qx sdk: sign failed err: %v", err)
			return nil, err
		}
	}
	res := qxTypes.QxClientApiSignResp{}
	_ = json.Unmarshal(result, &res)
	if res.Code != qxCodes.QxEngineStatusOK {
		if s.QxCtx.Cli.Config.AutoRetry {
			if s.QxCtx.Cli.RetryTimes > s.QxCtx.Cli.Config.MaxRetryTimes {
				s.Status = SdkStatusSignMaxTimes
				logx.Errorf("qx sdk: fail max times: %v", err)
				return nil, err
			} else {
				logx.Errorf("qx sdk: sign failed err: %v", err)
				s.QxCtx.Cli.RetryTimes++
				return s.AuthLogin()
			}
		} else {
			s.Status = SdkStatusLoginFailed
			logx.Infof("qx sdk: sign failed err: %v", res.Msg)
			return nil, err
		}
	}

	logx.Infof("qx sdk: api sign success")
	if s.QxCtx.Cli.Config.Debug {
		logx.Infof("qx sdk: api sign result: %s", string(result))
	}
	// note: 记录access token
	s.QxCtx.Cli.AccessToken = res.Data.AccessToken
	s.QxCtx.Cli.AccessTokenExpires = res.Data.AccessExpiresAt
	s.QxCtx.Cli.RefreshToken = res.Data.RefreshToken
	s.QxCtx.Cli.RefreshTokenExpires = res.Data.RefreshExpiresAt
	// note: 重置重试次数
	s.QxCtx.Cli.RetryTimes = 0
	if s.Status <= SdkStatusLogined {
		s.Status = SdkStatusLogined
	}
	return s, nil
}

func (s *QxSdk) AuthRefresh() (*QxSdk, error) {
	// note: 如果sdk未准备好，直接返回
	if s.Status < SdkStatusRemoteHealthy {
		logx.Errorf("qx sdk: not ready")
		return s, qxErrs.ErrNotReady
	}

	nowTime := time.Now()

	// note: 判断accessToken过期了没
	if (s.QxCtx.Cli.AccessTokenExpires - s.QxCtx.Cli.Config.Deadline) >= nowTime.Unix() {
		if s.QxCtx.Cli.Config.Debug {
			logx.Infof("qx sdk:accessToken没过期，过期时间为: %s, 当前时间为: %s", time.Unix(s.QxCtx.Cli.AccessTokenExpires, 0).Format(time.DateTime), nowTime.Format(time.DateTime))
		}
		// note: 没过期，直接返回
		return s, nil
	}
	if (s.QxCtx.Cli.RefreshTokenExpires - s.QxCtx.Cli.Config.Deadline) >= nowTime.Unix() {
		if s.QxCtx.Cli.Config.Debug {
			logx.Infof("qx sdk:accessToken过期了，过期时间为: %s, 但是refreshToken没过期，过期时间为: %s, 当前时间为: %s", time.Unix(s.QxCtx.Cli.AccessTokenExpires, 0).Format(time.DateTime), time.Unix(s.QxCtx.Cli.RefreshTokenExpires, 0).Format(time.DateTime), nowTime.Format(time.DateTime))
		}
		// note: refreshToken没过期，但是accessToken过期了
		reqFn := s.EasyUnLoginRequest(s.QxCtx.Cli.Context, "/auth/refresh", "POST", &qxTypes.QxClientApiRefreshReq{
			AccessKey:    s.QxCtx.Cli.Config.AccessKeyId,
			RefreshToken: s.QxCtx.Cli.RefreshToken,
		})
		result, err := reqFn()
		if err != nil {
			logx.Errorf("qx sdk: api refresh error: %v", err)
			if s.QxCtx.Cli.Config.AutoRetry {
				if s.QxCtx.Cli.RetryTimes > s.QxCtx.Cli.Config.MaxRetryTimes {
					s.Status = SdkStatusLoginFailed
					logx.Errorf("qx sdk: sdk fail max times: %v", err)
					return nil, err
				} else {
					logx.Errorf("qx sdk: qx sdk refresh failed, next try, err: %v, ", err)
					s.QxCtx.Cli.RetryTimes++
					return s.AuthRefresh()
				}
			} else {
				s.Status = SdkStatusLoginFailed
				logx.Infof("qx sdk: refresh failed err: %v", err)
				return nil, err
			}
		}
		res := qxTypes.QxClientApiRefreshResp{}
		_ = json.Unmarshal(result, &res)
		if res.Code != qxCodes.QxEngineStatusOK {
			if s.QxCtx.Cli.Config.AutoRetry {
				if s.QxCtx.Cli.RetryTimes > s.QxCtx.Cli.Config.MaxRetryTimes {
					s.Status = SdkStatusSignMaxTimes
					logx.Errorf("qx sdk: fail max times: %v", err)
					return nil, err
				} else {
					logx.Errorf("qx sdk: refresh failed err: %v", err)
					s.QxCtx.Cli.RetryTimes++
					return s.AuthRefresh()
				}

			} else {
				s.Status = SdkStatusLoginFailed
				logx.Infof("qx sdk: refresh failed err: %v", res.Msg)
				return nil, err
			}
		}

		if s.QxCtx.Cli.Config.Debug {
			logx.Infof("qx sdk: refresh result: %s", string(result))
		}
		// note: 记录access token
		s.QxCtx.Cli.AccessToken = res.Data.AccessToken
		s.QxCtx.Cli.AccessTokenExpires = res.Data.AccessExpiresAt
		// note: 重置重试次数
		s.QxCtx.Cli.RetryTimes = 0
		if s.Status <= SdkStatusLogined {
			s.Status = SdkStatusLogined
		}
	} else {
		// note: refreshToken过期了
		logx.Errorf("qx sdk: refreshToken 过期了")
		if s.Status <= SdkStatusRemoteHealthy {
			s.Status = SdkStatusRemoteHealthy
		}
		return s.AuthLogin()
	}
	return s, nil
}

func (s *QxSdk) KeyExChange() (*QxSdk, error) {
	// note: 如果sdk未准备好，直接返回
	if s.Status < SdkStatusLogined {
		logx.Errorf("qx sdk: sdk not logined")
		return s, qxErrs.ErrNotLogined
	}

	if s.QxCtx.Parser.Status() == qxParser.QxParserStatusReady {
		nowTime := time.Now()
		// note: 判断客户端证书过期了没
		if (s.QxCtx.Parser.ExpireAt().Unix() - s.QxCtx.Cli.Config.Deadline) >= nowTime.Unix() {
			if s.QxCtx.Cli.Config.Debug {
				logx.Infof("qx sdk: 客户端证书没过期，过期时间为: %s, 当前时间为: %s", s.QxCtx.Parser.ExpireAt().Format(time.DateTime), nowTime.Format(time.DateTime))
			}
			// note: 没过期，直接返回
			return s, nil
		}
		if s.QxCtx.Cli.Config.Debug {
			logx.Infof("qx sdk: 客户端证书过期了，过期时间为: %s", s.QxCtx.Parser.ExpireAt().Format(time.DateTime))
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Status <= SdkStatusStartKeyExChange {
		s.Status = SdkStatusStartKeyExChange
	}
	// note: 客户端生成一张证书
	privKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		logx.Errorf("qx sdk: sdk generate key err: %v", err)
		return s, err
	}
	if s.Status <= SdkStatusKeyExChanging {
		s.Status = SdkStatusKeyExChanging
	}
	// note: 发送给网关交换公钥
	tmpPub := base64.StdEncoding.EncodeToString(privKey.PublicKey().Bytes())
	reqFn := s.QxCtx.Cli.EasyNewRequest(s.QxCtx.Cli.Context, "/keyExchange", http.MethodPost, &qxTypes.QxClientKeyExChangeReq{
		AccessKey: s.QxCtx.Cli.AccessKeyId,
		PublicKey: tmpPub,
	})
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk: keyexchange request error: %v", err)
		return s, err
	}
	result := qxTypes.QxClientKeyExChangeResp{}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: keyexchange  fail: %v", result)
		return s, errors.New(result.Msg)
	}
	if s.QxCtx.Cli.Config.Debug {
		logx.Infof("qx sdk: keyexchange request success: %v", result)
	}

	// note: 完成解析器的初始化
	tmpPubkey, err := base64.StdEncoding.DecodeString(result.Data.PublicKey)
	if err != nil {
		logx.Errorf("qx sdk: keyexchange load tmpPubkey error: %v", err)
		return nil, err
	}
	// 解析远端公钥
	remotePubKey, err := ecdh.P256().NewPublicKey(tmpPubkey)
	if err != nil {
		logx.Errorf("qx sdk: keyexchange parse remote public key error: %v", err)
		return nil, err
	}
	expireAt := time.Unix(result.Data.ExpireAt, 0)
	if err = s.QxCtx.Parser.Init(privKey, remotePubKey, expireAt); err != nil {
		logx.Errorf("qx sdk: qx sdk parser init err: %v", err)
		return nil, err
	}
	s.QxCtx.Cli.SetSessionId(result.Data.SessionId)
	if s.Status <= SdkStatusReady {
		s.Status = SdkStatusReady
	}

	return s, nil
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
	s.QxCtx.Cli.Client.CloseIdleConnections()

	fmt.Println("SDK resources released.")
}

func (s *QxSdk) GetVersion() string {
	return s.Version
}

func (s *QxSdk) GetStatus() int {
	return s.Status
}

func (s *QxSdk) FormatQxSdkStatus() string {
	switch s.Status {
	case SdkStatusNotReady:
		return "未就绪"
	case SdkStatusRemoteUnHealthy:
		return "无法连接服务器"
	case SdkStatusRemoteHealthy:
		return "远程服务器健康检查已通过"
	case SdkStatusSignMaxTimes:
		return "登录次数超限"
	case SdkStatusLoginFailed:
		return "登录失败"
	case SdkStatusLogined:
		return "已登录"
	case SdkStatusStartKeyExChange:
		return "开始密钥交换"
	case SdkStatusKeyExChanging:
		return "密钥交换中"
	case SdkStatusReady:
		return "已就绪"
	}
	return "未知状态"
}
