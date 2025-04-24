package qxCli

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxConfig"
	"github.com/Technology-99/qxLib/qxCommonHeader"
	"github.com/Technology-99/qxLib/qxMiddleware"
	"github.com/Technology-99/qxLib/qxSony"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"time"
)

type QxClient struct {
	*http.Client
	Config              *qxConfig.Config
	Context             context.Context
	SessionId           string
	AccessKeyId         string
	AccessToken         string
	AccessTokenExpires  int64
	RefreshToken        string
	RefreshTokenExpires int64
	RetryTimes          int
}

func NewQxClient(ctx context.Context, conf *qxConfig.Config) *QxClient {
	httpClient := &http.Client{
		Timeout: qxConfig.DefaultTimeout * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	cli := &QxClient{
		Config:      conf,
		AccessKeyId: conf.AccessKeyId,
		Client:      httpClient,
		Context:     ctx,
		//Status:      StatusNotReady,
	}
	return cli
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

func (cli *QxClient) SetSessionId(sessionId string) *QxClient {
	cli.SessionId = sessionId
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

func (cli *QxClient) EasyNewRequest(ctx context.Context, relativePath string, method string, sendBody interface{}) func() ([]byte, error) {
	apiUrl := fmt.Sprintf("%s://%s%s%s", cli.Config.Protocol, cli.Config.Endpoint, "/qx/v5/apis", relativePath)
	if cli.Context.Value(qxMiddleware.CtxRequestID) != nil {
		logx.Infof("requestID: %s, EasyNewRequest url: %s", cli.Context.Value(qxMiddleware.CtxRequestID), apiUrl)
	} else {
		logx.Infof("EasyNewRequest url: %s", apiUrl)
	}
	return cli.NewRequest(ctx, apiUrl, method, cli.GenHeaders(), sendBody)
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

	headers[qxCommonHeader.HeaderAuthorization] = "Bearer " + cli.AccessToken
	headers[qxCommonHeader.HeaderXAccessKeyFor] = cli.AccessKeyId
	headers[qxCommonHeader.HeaderXSessionIdFor] = cli.SessionId
	return &headers
}
