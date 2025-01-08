package sdk

import (
	"bytes"
	"context"
	"crypto/tls"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/req"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/resp"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/types"
	"github.com/Technology-99/third_party/commKey"
	"github.com/Technology-99/third_party/middleware"
	"github.com/Technology-99/third_party/response"
	"github.com/Technology-99/third_party/sony"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultTimeout = 2000
	keyRequestId   = "requestId"
)

//go:embed "version"
var VersionF embed.FS

type Sdk struct {
	Version string
	Config  *Config

	Context context.Context
	Client  *http.Client

	Status              int
	AccessToken         string
	AccessTokenExpires  int64
	RefreshToken        string
	RefreshTokenExpires int64
	RetryTimes          int

	//Auth         *Auth
	//CloudC       *CloudC
	//ThirdParty   *ThirdParty
	//Sls          *Sls
	//Mix          *Mix
	//H5           *H5
	//Transactions *Transactions
	//AIot         *AIot
}

func NewSdk(AccessKeyId, AccessKeySecret, Endpoint string) *Sdk {
	ctx := context.Background()
	c := DefaultConfig(AccessKeyId, AccessKeySecret, Endpoint)

	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{
		Timeout: defaultTimeout * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return &Sdk{
		Version: string(versionFile),
		Config:  c,
		Status:  types.STATUS_NOT_READY,
		Context: ctx,
		Client:  httpClient,
	}
}
func (s *Sdk) GetVersion() string {
	return s.Version
}

func (s *Sdk) WithConfig(config *Config) *Sdk {
	s.Config = config
	return s
}

func (s *Sdk) WithContext(ctx context.Context) *Sdk {
	s.Context = ctx
	return s
}

// note: 添加将requestID继承到下个服务的能力
func (s *Sdk) WithRequestId(requestId string) *Sdk {
	s.Context = context.WithValue(s.Context, middleware.CtxRequestID, requestId)
	return s
}

func (s *Sdk) WithTimeout(timeout time.Duration) *Sdk {
	s.Config.Timeout = timeout
	s.Client.Timeout = timeout
	return s
}

func (s *Sdk) NewRequest(Url string, method string, sendBody interface{}) func() ([]byte, error) {
	var body []byte
	var err error

	// note: 先处理请求头
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	// note: 先判断ctx上是否存在requestId
	if value := s.Context.Value(keyRequestId); value != nil {
		headers[commKey.HeaderXRequestIDFor] = value.(string)
	} else {
		headers[commKey.HeaderXRequestIDFor] = sony.NextId()
	}
	// note: 再判断是否登录成功
	if s.Status == types.STATUS_LOGINED {
		headers[commKey.HeaderAuthorization] = "Bearer " + s.AccessToken
		headers[commKey.HeaderXAccessKeyFor] = s.Config.AccessKeyId
	}

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)

		sendBodyBt, err := json.Marshal(sendBody)
		if err != nil {
			return
		}

		req, err := http.NewRequest(method, Url, bytes.NewBuffer(sendBodyBt))
		if err != nil {
			logx.Errorf("requestID: %s, Error creating request: %v", err)
			return
		}
		for key, val := range headers {
			req.Header.Set(key, val)
		}

		var res *http.Response
		res, err = s.Client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}

		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
		return
	}()
	return func() ([]byte, error) {
		_, ok := <-c
		if !ok {
			//fmt.Println("channel closed!")
			return body, err
		}
		return body, err
	}
}

func (s *Sdk) AutoAuth() *Sdk {
	s, _ = s.AuthHealthZ().AuthLogin()
	go s.AutoRefresh()
	return s
}

func (s *Sdk) AuthHealthZ() *Sdk {
	apiUrl := fmt.Sprintf("%s://%s/qiongxiao/v5/apis/healthz", s.Config.Protocol, s.Config.Endpoint)
	reqFn := s.NewRequest(apiUrl, "GET", nil)
	result, err := reqFn()
	if err != nil {
		logx.Errorf("healthz request error: %v", err)
		return nil
	}
	res := resp.DefaultResponseBase{}
	_ = json.Unmarshal(result, &res)
	if res.Code == response.SUCCESS {
		logx.Infof("sdk healthz success")
		s.Status = types.STATUS_READY
	} else {
		panic(res.Msg)
	}
	return s
}

func (s *Sdk) AuthLogin() (*Sdk, error) {
	logx.Infof("打印sdk的状态: %s", s.FormatSdkStatus())
	if s.Status == types.STATUS_NOT_READY {
		logx.Errorf("sdk not ready")
		return s, types.ErrNotReady
	}

	apiUrl := fmt.Sprintf("%s://%s/qiongxiao/v5/apis/auth/api/sign", s.Config.Protocol, s.Config.Endpoint)
	reqFn := s.NewRequest(apiUrl, "POST", &req.QxV5ApisApiSignReq{
		AccessKey:    s.Config.AccessKeyId,
		AccessSecret: s.Config.AccessKeySecret,
	})
	result, err := reqFn()
	if err != nil {
		logx.Errorf("api sign error: %v", err)
		if s.Config.AutoRetry {
			if s.RetryTimes > s.Config.MaxRetryTimes {
				s.Status = types.STATUS_NOT_READY
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
		if s.Config.Debug {
			logx.Infof("sdk api sign result: %s", string(result))
		}
		// note: 记录access token
		s.AccessToken = res.Data.AccessToken
		s.AccessTokenExpires = res.Data.AccessExpiresAt
		s.RefreshToken = res.Data.RefreshToken
		s.RefreshTokenExpires = res.Data.RefreshExpiresAt
		s.AuthSuccess()
	} else {
		if s.Config.AutoRetry {
			if s.RetryTimes > s.Config.MaxRetryTimes {
				s.Status = types.STATUS_NOT_READY
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
	if s.Config.Debug {
		logx.Infof("打印sdk的状态: %s", s.FormatSdkStatus())
	}
	// note: 如果链接未准备好，直接返回
	if s.Status == types.STATUS_NOT_READY {
		return nil, types.ErrNotReady
	}

	nowTime := time.Now()

	// note: 判断accessToken过期了没
	if (s.AccessTokenExpires - s.Config.Deadline) >= nowTime.Unix() {
		if s.Config.Debug {
			logx.Infof("accessToken没过期，过期时间为: %s, 当前时间为: %s", time.Unix(s.AccessTokenExpires, 0).Format(time.DateTime), nowTime.Format(time.DateTime))
		}
		// note: 没过期，直接返回
		return s, nil
	}
	if (s.RefreshTokenExpires - s.Config.Deadline) >= nowTime.Unix() {
		if s.Config.Debug {
			logx.Infof("accessToken过期了，过期时间为: %s, 但是refreshToken没过期，过期时间为: %s, 当前时间为: %s", time.Unix(s.AccessTokenExpires, 0).Format(time.DateTime), time.Unix(s.RefreshTokenExpires, 0).Format(time.DateTime), nowTime.Format(time.DateTime))
		}
		// note: refreshToken没过期，但是accessToken过期了
		apiUrl := fmt.Sprintf("%s://%s/qiongxiao/v5/apis/auth/api/refresh", s.Config.Protocol, s.Config.Endpoint)
		reqFn := s.NewRequest(apiUrl, "POST", &req.QxV5ApisApiRefreshReq{
			AccessKey:    s.Config.AccessKeyId,
			RefreshToken: s.RefreshToken,
		})
		result, err := reqFn()
		if err != nil {
			logx.Errorf("api refresh error: %v", err)
			if s.Config.AutoRetry {
				if s.RetryTimes > s.Config.MaxRetryTimes {
					s.Status = types.STATUS_NOT_READY
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
			if s.Config.Debug {
				logx.Infof("sdk api refresh result: %s", string(result))
			}
			// note: 记录access token
			s.AccessToken = res.Data.AccessToken
			s.AccessTokenExpires = res.Data.AccessExpiresAt
			s.AuthSuccess()
		} else {
			if s.Config.AutoRetry {
				if s.RetryTimes > s.Config.MaxRetryTimes {
					s.Status = types.STATUS_NOT_READY
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
		s.Status = types.STATUS_NOT_READY
		return s.AuthHealthZ().AuthLogin()
	}
	return s, nil
}

func (s *Sdk) AutoRefresh() *Sdk {
	if s.Config.AutoRefreshToken {
		// note: check refresh token is expired
		for {
			if s.RetryTimes > s.Config.MaxRetryTimes {
				// note: close auto refresh
				s.Config.AutoRefreshToken = false
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
			time.Sleep(time.Second)
		}
	}
	return s
}

func (s *Sdk) AuthSuccess() {
	s.RetryTimes = 0
	s.Status = types.STATUS_LOGINED
}

func (s *Sdk) AuthFail(err error) {
	if s.Config.AutoRetry {
		s.RetryTimes++
	} else {
		s.Status = types.STATUS_NOT_READY
		panic(err)
	}
}

func (s *Sdk) FormatSdkStatus() string {
	switch s.Status {
	case types.STATUS_READY:
		return "已就绪"
	case types.STATUS_LOGINED:
		return "已登录"
	case types.STATUS_NOT_READY:
		return "未就绪"
	}
	return "未知状态"
}

func (s *Sdk) SonyCtx() context.Context {
	requestID := ""
	if value := s.Context.Value(keyRequestId); value != nil {
		requestID = value.(string)
	} else {
		requestID = sony.NextId()
	}
	logx.Debugf("requestId: %s", requestID)
	md := metadata.New(map[string]string{
		"X-RequestID-For": requestID,
		"X-AccessKey-For": s.Config.AccessKeyId,
		"Authorization":   "Bearer " + s.AccessToken,
	})
	s.Context = metadata.NewOutgoingContext(s.Context, md)
	return s.Context
}
