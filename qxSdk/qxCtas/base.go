package qxCtas

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
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
		qxCtx *qxCtx.QxCtx
	}
)

func NewCtasBaseService(qxCtx *qxCtx.QxCtx) CtasBaseService {
	return &defaultCtasBaseService{
		qxCtx: qxCtx,
	}
}

func (m *defaultCtasBaseService) TestMsg(ctx context.Context, params *qxTypesCtas.CtasTestMsgReq) (result *qxTypesCtas.CtasTestMsgResp, err error) {
	result = &qxTypesCtas.CtasTestMsgResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/captcha/generate", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}

	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}
