package qxSdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxCommonHeader"
	"github.com/Technology-99/qxLib/qxMiddleware"
	"github.com/Technology-99/qxLib/qxSony"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// note: 解析器测试请求
func (s *QxSdk) TestMsg() *QxSdk {
	msg := fmt.Sprintf("test msg: %s", time.Now().Format(time.DateTime))
	encryptMsg, err := s.QxCtx.Parser.Encrypt(msg)
	if err != nil {
		logx.Errorf("qx sdk encrypt error: %v", err)
		return nil
	}
	if s.QxCtx.Cli.Config.Debug {
		logx.Infof("qx sdk encrypt msg: %s", encryptMsg)
	}
	reqFn := s.QxCtx.Cli.EasyNewRequest(s.QxCtx.Cli.Context, "/test", "POST", qxTypes.QxClientTestMsgReq{
		Msg: encryptMsg,
	})
	result, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk testMsg request error: %v", err)
		return nil
	}
	if s.QxCtx.Cli.Config.Debug {
		logx.Infof("qx sdk testMsg result: %s", string(result))
	}
	res := qxTypes.QxClientTestMsgResp{}
	_ = json.Unmarshal(result, &res)
	if res.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk testMsg request error: %v", res.Msg)
	}
	if s.QxCtx.Cli.Config.Debug {
		logx.Infof("qx sdk testMsg res: %v", res)
	}
	decryptMsg, err := s.QxCtx.Parser.Decrypt(res.Data)
	if err != nil {
		logx.Errorf("qx sdk encrypt error: %v", err)
		return nil
	}

	if s.QxCtx.Cli.Config.Debug {
		logx.Infof("qx sdk testMsg decrypt data %s", string(decryptMsg))
	}
	logx.Infof("qx sdk encrypt msg: %s", decryptMsg)
	return s
}

// note: 自动验证远端健康
func (s *QxSdk) AutoHealthZ() *QxSdk {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done() // 确保任务退出时通知 WaitGroup
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.HealthZ()
			case <-s.ctx.Done():
				fmt.Println("healthz stopped.")
				return
			}
		}
	}()
	return s
}

func (s *QxSdk) AutoRefresh() *QxSdk {
	if s.QxCtx.Cli.Config.AutoRefreshToken {
		s.wg.Add(1)

		go func() {
			defer s.wg.Done() // 确保任务退出时通知 WaitGroup
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					if s.QxCtx.Cli.Config.Debug {
						logx.Debugf("qx sdk auto refresh")
					}
					if _, err := s.AuthRefresh(); err != nil {
						logx.Errorf("qx sdk auto refresh fail: %+v", err)
						time.Sleep(time.Second)
						continue
					}
				case <-s.ctx.Done():
					fmt.Println("qx sdk auto refresh stopped.")
					return
				}
			}
		}()
	}
	return s
}

func (s *QxSdk) AutoKeyExChange() *QxSdk {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done() // 确保任务退出时通知 WaitGroup
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if s.QxCtx.Cli.Config.Debug {
					logx.Debugf("qx sdk auto key exchange")
				}
				if _, err := s.KeyExChange(); err != nil {
					logx.Errorf("qx sdk auto key exchange fail: %+v", err)
					time.Sleep(time.Second)
					continue
				}
			case <-s.ctx.Done():
				fmt.Println("qx sdk auto refresh stopped.")
				return
			}
		}
	}()
	return s
}

func (s *QxSdk) EasyUnLoginRequest(ctx context.Context, relativePath string, method string, sendBody interface{}) func() ([]byte, error) {
	apiUrl := fmt.Sprintf("%s://%s%s%s", s.QxCtx.Cli.Config.Protocol, s.QxCtx.Cli.Config.Endpoint, "/qx/v5/apis", relativePath)
	return s.QxCtx.Cli.NewRequest(ctx, apiUrl, method, s.GenUnLoginHeaders(), sendBody)
}

func (s *QxSdk) GenUnLoginHeaders() *map[string]string {
	// note: 先处理请求头
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	if value := s.QxCtx.Cli.Context.Value(qxMiddleware.CtxRequestID); value != nil {
		headers[qxCommonHeader.HeaderXRequestIDFor] = value.(string)
	} else {
		headers[qxCommonHeader.HeaderXRequestIDFor] = qxSony.NextId()
	}
	return &headers
}
