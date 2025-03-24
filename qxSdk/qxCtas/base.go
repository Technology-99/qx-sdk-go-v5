package qxCtas

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesCtas"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	CtasBaseService interface {
		TestMsg(ctx context.Context, params *qxTypesCtas.CtasTestMsgReq) (result *qxTypesCtas.CtasTestMsgResp, err error)
	}
	defaultCtasBaseService struct {
		cli *qxCli.QxClient
	}
)

func NewCtasBaseService(cli *qxCli.QxClient) CtasBaseService {
	return &defaultCtasBaseService{
		cli: cli,
	}
}

func (m *defaultCtasBaseService) TestMsg(ctx context.Context, params *qxTypesCtas.CtasTestMsgReq) (result *qxTypesCtas.CtasTestMsgResp, err error) {
	result = &qxTypesCtas.CtasTestMsgResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/captcha/generate", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("healthz request error: %v", err)
		return nil, nil
	}

	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qiongxiao sdk errlog: captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}
