package qxMas

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	MasBaseService interface {
		// note: 生成验证码
		CaptchaGenerate(ctx context.Context, params *qxTypes.ApiCaptchaGenerateReq) (result *qxTypes.ApiCaptchaGenerateResp, err error)
		SmsSend(ctx context.Context, params *qxTypes.ApiSmsSendReq) (result *qxTypes.ApiSmsSendResp, err error)
	}

	defaultMasBaseService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewMsgBaseService(qxCtx *qxCtx.QxCtx) MasBaseService {
	return &defaultMasBaseService{
		qxCtx: qxCtx,
	}
}

func (m *defaultMasBaseService) CaptchaGenerate(ctx context.Context, params *qxTypes.ApiCaptchaGenerateReq) (result *qxTypes.ApiCaptchaGenerateResp, err error) {
	result = &qxTypes.ApiCaptchaGenerateResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/mas/captcha/generate", http.MethodPost, &params)
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

func (m *defaultMasBaseService) SmsSend(ctx context.Context, params *qxTypes.ApiSmsSendReq) (result *qxTypes.ApiSmsSendResp, err error) {
	result = &qxTypes.ApiSmsSendResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/mas/sms/send", http.MethodPost, &params)
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
