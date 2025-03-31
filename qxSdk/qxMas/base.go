package qxMas

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesMas"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	MasBaseService interface {
		// note: 生成验证码
		CaptchaGenerate(ctx context.Context, params *qxTypes.ApiCaptchaGenerateReq) (result *qxTypes.ApiCaptchaGenerateResp, err error)
		SmsSend(ctx context.Context, params *qxTypes.ApiSmsSendReq) (result *qxTypes.ApiSmsSendResp, err error)

		BehavioralVerificationInit(ctx context.Context, params *qxTypesMas.BehavioralVerificationInitReq) (result *qxTypesMas.BehavioralVerificationInitResp, err error)
		BehavioralVerificationVerify(ctx context.Context, params *qxTypesMas.BehavioralVerificationVerifyReq) (result *qxTypesMas.BehavioralVerificationVerifyResp, err error)
		SmsVerificationInit(ctx context.Context, params *qxTypesMas.SmsInitReq) (result *qxTypesMas.SmsInitResp, err error)
		SmsVerificationVerify(ctx context.Context, params *qxTypesMas.SmsVerifyReq) (result *qxTypesMas.SmsVerifyResp, err error)
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

func (m *defaultMasBaseService) BehavioralVerificationInit(ctx context.Context, params *qxTypesMas.BehavioralVerificationInitReq) (result *qxTypesMas.BehavioralVerificationInitResp, err error) {
	result = &qxTypesMas.BehavioralVerificationInitResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/mas/bv/init", http.MethodPost, &params)
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

func (m *defaultMasBaseService) BehavioralVerificationVerify(ctx context.Context, params *qxTypesMas.BehavioralVerificationVerifyReq) (result *qxTypesMas.BehavioralVerificationVerifyResp, err error) {
	result = &qxTypesMas.BehavioralVerificationVerifyResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/mas/bv/verify", http.MethodPost, &params)
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

func (m *defaultMasBaseService) SmsVerificationInit(ctx context.Context, params *qxTypesMas.SmsInitReq) (result *qxTypesMas.SmsInitResp, err error) {
	result = &qxTypesMas.SmsInitResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/mas/sms/init", http.MethodPost, &params)
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

func (m *defaultMasBaseService) SmsVerificationVerify(ctx context.Context, params *qxTypesMas.SmsVerifyReq) (result *qxTypesMas.SmsVerifyResp, err error) {
	result = &qxTypesMas.SmsVerifyResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/mas/sms/verify", http.MethodPost, &params)
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
