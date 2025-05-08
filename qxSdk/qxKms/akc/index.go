package akc

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesKms"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	KmsAkcService interface {
		KmsAkcCreateKeychain(ctx context.Context, params *qxTypesKms.KmsAkcCreateKeychainReq) (result *qxTypesKms.KmsAkcCreateKeychainResp, err error)
		KmsAkcSign(ctx context.Context, params *qxTypesKms.KmsAkcSignReq) (result *qxTypesKms.KmsAkcSignResp, err error)
		KmsAkcVerify(ctx context.Context, params *qxTypesKms.KmsAkcVerifyReq) (result *qxTypesKms.KmsAkcVerifyResp, err error)
	}

	defaultKmsAkcService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewKmsAkcService(qxCtx *qxCtx.QxCtx) KmsAkcService {
	// note: 初始化Kms系统
	return &defaultKmsAkcService{
		qxCtx: qxCtx,
	}
}

func (m *defaultKmsAkcService) KmsAkcCreateKeychain(ctx context.Context, params *qxTypesKms.KmsAkcCreateKeychainReq) (result *qxTypesKms.KmsAkcCreateKeychainResp, err error) {
	result = &qxTypesKms.KmsAkcCreateKeychainResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/akc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsAkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsAkcService) KmsAkcSign(ctx context.Context, params *qxTypesKms.KmsAkcSignReq) (result *qxTypesKms.KmsAkcSignResp, err error) {
	result = &qxTypesKms.KmsAkcSignResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/akc/sign", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsAkcSign fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsAkcService) KmsAkcVerify(ctx context.Context, params *qxTypesKms.KmsAkcVerifyReq) (result *qxTypesKms.KmsAkcVerifyResp, err error) {
	result = &qxTypesKms.KmsAkcVerifyResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/akc/verify", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsAkcVerify fail: %v", result)
		return result, nil
	}
	return result, nil
}
