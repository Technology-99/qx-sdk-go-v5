package qxCli

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxConfig"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxParser"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxCommonHeader"
	"github.com/Technology-99/qxLib/qxMiddleware"
	"github.com/Technology-99/qxLib/qxRes"
	"github.com/Technology-99/qxLib/qxSony"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type QxClient struct {
	*http.Client

	Config *qxConfig.Config
	// note: 锁
	mu             sync.Mutex
	Context        context.Context
	RetryTimes     int
	Parser         qxParser.QxParser
	authConfig     qxTypes.AuthConfig
	authConfigFile string
	viper          *viper.Viper
}

func NewQxClient(ctx context.Context, conf *qxConfig.Config) *QxClient {
	httpClient := &http.Client{
		Timeout: qxConfig.DefaultTimeout * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	authConfigFile := path.Join(conf.HomeDir, ".qxConfig/config.json")
	cli := &QxClient{
		Config:         conf,
		Client:         httpClient,
		Context:        ctx,
		Parser:         qxParser.NewQxParser(),
		authConfigFile: authConfigFile,
		viper:          viper.New(),
	}

	cli.Init()
	return cli
}

func (cli *QxClient) Init() error {
	cli.viper.SetConfigFile(cli.authConfigFile)
	cli.viper.SetConfigType("json")
	if err := cli.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logx.Errorf("文件不存在")
			// Config file not found; ignore error if desired
			if err = cli.Login(); err != nil {
				log.Fatalf("qx sdk: login err: %v", err)
			}
		} else {
			// Config file was found but another error was produced
			logx.Errorf("qx sdk: read config err: %v", err)
			os.MkdirAll(path.Dir(cli.authConfigFile), os.ModePerm)
			if err = cli.Login(); err != nil {
				log.Fatalf("qx sdk: login err: %v", err)
			}
		}
	}
	logx.Infof("qx sdk: read config file success")

	// 监听文件变化
	cli.viper.WatchConfig()
	cli.viper.OnConfigChange(func(e fsnotify.Event) {
		logx.Infof("qx sdk: Config file changed: %s", e.Name)
	})
	return nil
}

func (cli *QxClient) Login() error {
	result, err := cli.EasyNewRequest(cli.Context, "/auth/token", "POST", &qxTypes.QxClientApiTokenReq{
		AccessKey:    cli.Config.AccessKeyId,
		AccessSecret: cli.Config.AccessKeySecret,
	})
	if err != nil {
		if cli.Config.AutoRetry {
			if cli.RetryTimes > cli.Config.MaxRetryTimes {
				logx.Errorf("qx sdk: fail max times: %v", err)
				return err
			} else {
				logx.Errorf("qx sdk: token failed, next try, err: %v, ", err)
				cli.RetryTimes++
				return cli.Login()
			}
		} else {
			logx.Infof("qx sdk: token failed err: %v", err)
			return err
		}
	}
	res := qxTypes.QxClientApiTokenResp{}
	_ = json.Unmarshal(result, &res)
	if res.Code != qxCodes.QxEngineStatusOK {
		if cli.Config.AutoRetry {
			if cli.RetryTimes > cli.Config.MaxRetryTimes {
				logx.Errorf("qx sdk: fail max times: %v", err)
				return err
			} else {
				logx.Errorf("qx sdk: token failed, next try, err: %v, ", err)
				cli.RetryTimes++
				return cli.Login()
			}
		} else {
			logx.Infof("qx sdk: sign failed err: %v", res.Msg)
			return errors.New(res.Msg)
		}
	}

	cli.viper.Set(fmt.Sprintf("auths.%s.accessToken", cli.Config.AccessKeyId), res.Data.AccessToken)
	cli.viper.Set(fmt.Sprintf("auths.%s.expiresIn", cli.Config.AccessKeyId), res.Data.ExpiresIn)

	cli.viper.WriteConfig()

	logx.Infof("qx sdk: api token success: %s", res.Data.AccessToken)

	return nil
}

func (cli *QxClient) WithContext(ctx context.Context) *QxClient {
	cli.Context = ctx
	return cli
}

// note: 添加将requestID继承到下个服务的能力
func (cli *QxClient) WithRequestId(requestId string) *QxClient {
	cli.Context = context.WithValue(cli.Context, qxMiddleware.CtxRequestID, requestId)
	return cli
}

func (cli *QxClient) WithTimeout(timeout time.Duration) *QxClient {
	cli.Client.Timeout = timeout
	return cli
}

func (cli *QxClient) NewRequest(
	ctx context.Context, // 新增 context 参数
	url string, // URL
	method string, // HTTP 方法
	headers *map[string]string, // 请求头
	sendBody interface{}) func() ([]byte, error) { // 返回闭包函数

	var (
		body []byte
		err  error
	)

	// 创建一个 channel 来控制请求完成或超时
	c := make(chan struct{})
	go func() {
		defer close(c) // 保证 goroutine 退出时关闭 channel

		sendBodyJson := ""

		if sendBody != nil {
			// 将发送体序列化为 JSON
			sendBodyBt, marshalErr := json.Marshal(sendBody)
			if marshalErr != nil {
				err = marshalErr
				return
			}
			sendBodyJson = string(sendBodyBt)
		}

		// 使用 context 控制请求
		req, reqErr := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer([]byte(sendBodyJson)))
		if reqErr != nil {
			err = reqErr
			return
		}

		// 设置请求头
		if headers != nil {
			for k, v := range *headers {
				req.Header.Set(k, v)
			}
		}

		// 发送请求
		var res *http.Response
		res, err = cli.Client.Do(req)
		if err != nil {
			return
		}
		defer res.Body.Close()

		// 读取响应体
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}

	}()

	return func() ([]byte, error) {
		select {
		case <-c: // 请求完成
			return body, err
		case <-ctx.Done(): // 请求超时或取消
			return nil, ctx.Err()
		}
	}
}

func (cli *QxClient) EasyNewRequest(ctx context.Context, relativePath string, method string, sendBody interface{}) ([]byte, error) {
	apiUrl := fmt.Sprintf("%s://%s%s%s", cli.Config.Protocol, cli.Config.Endpoint, "/qx/v5/apis", relativePath)
	if cli.Context.Value(qxMiddleware.CtxRequestID) != nil {
		logx.Infof("requestID: %s, EasyNewRequest url: %s", cli.Context.Value(qxMiddleware.CtxRequestID), apiUrl)
	} else {
		logx.Infof("EasyNewRequest url: %s", apiUrl)
	}
reReq:
	fn := cli.NewRequest(ctx, apiUrl, method, cli.GenHeaders(), sendBody)
	body, err := fn()
	// 解析响应体
	result := qxRes.BaseResponse[any]{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	logx.Infof("EasyNewRequest result: %v", result)
	if result.Code == qxCodes.QxEngineStatusAccessExpired || result.Code == qxCodes.QxEngineStatusAccessTokenInvalid || result.Code == qxCodes.QxEngineStatusTokenNotActiveYet {
		if err = cli.Login(); err != nil {
			logx.Errorf("qx sdk: init error: %v", err)
			return nil, err
		}
		logx.Infof("EasyNewRequest access token expired, retrying...")
		time.Sleep(time.Second * 5)
		logx.Infof("EasyNewRequest access token expired, retry req")
		goto reReq
	}
	return body, nil
}

func (cli *QxClient) GenHeaders() *map[string]string {
	// note: 先处理请求头
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	//headers[qxCommonHeader.HeaderXAuthMethodFor] = "api"
	// note: 先判断ctx上是否存在requestId
	if value := cli.Context.Value(qxMiddleware.CtxRequestID); value != nil {
		headers[qxCommonHeader.HeaderXRequestIDFor] = value.(string)
	} else {
		headers[qxCommonHeader.HeaderXRequestIDFor] = qxSony.NextId()
	}
	acsTokenKey := fmt.Sprintf("auths.%s.accessToken", cli.Config.AccessKeyId)
	acsToken := fmt.Sprintf("%s", cli.viper.Get(acsTokenKey))

	headers[qxCommonHeader.HeaderAuthorization] = "Bearer " + acsToken
	return &headers
}
