package skc

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
	KmsSkcService interface {
		KmsSkcCreateKeychain(ctx context.Context, params *qxTypesKms.KmsSkcCreateKeychainReq) (result *qxTypesKms.KmsSkcCreateKeychainResp, err error)
		KmsSkcEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcEncryptReq) (result *qxTypesKms.KmsSkcEncryptResp, err error)
		KmsSkcDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcDecryptReq) (result *qxTypesKms.KmsSkcDecryptResp, err error)
		KmsSkcBatchEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchEncryptReq) (result *qxTypesKms.KmsSkcBatchEncryptResp, err error)
		KmsSkcBatchDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchDecryptReq) (result *qxTypesKms.KmsSkcBatchDecryptResp, err error)
		KmsSkcCompare(ctx context.Context, params *qxTypesKms.KmsSkcCompareReq) (result *qxTypesKms.KmsSkcCompareResp, err error)
	}

	defaultKmsSkcService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewKmsSkcService(qxCtx *qxCtx.QxCtx) KmsSkcService {
	// note: 初始化Kms系统
	return &defaultKmsSkcService{
		qxCtx: qxCtx,
	}
}

func (m *defaultKmsSkcService) KmsSkcCreateKeychain(ctx context.Context, params *qxTypesKms.KmsSkcCreateKeychainReq) (result *qxTypesKms.KmsSkcCreateKeychainResp, err error) {
	result = &qxTypesKms.KmsSkcCreateKeychainResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/createKeychain", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcEncryptReq) (result *qxTypesKms.KmsSkcEncryptResp, err error) {
	result = &qxTypesKms.KmsSkcEncryptResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/encrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcDecryptReq) (result *qxTypesKms.KmsSkcDecryptResp, err error) {
	result = &qxTypesKms.KmsSkcDecryptResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/decrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcDecrypt fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchEncryptReq) (result *qxTypesKms.KmsSkcBatchEncryptResp, err error) {
	result = &qxTypesKms.KmsSkcBatchEncryptResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/batchEncrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchDecryptReq) (result *qxTypesKms.KmsSkcBatchDecryptResp, err error) {
	result = &qxTypesKms.KmsSkcBatchDecryptResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/batchDecrypt", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcCompare(ctx context.Context, params *qxTypesKms.KmsSkcCompareReq) (result *qxTypesKms.KmsSkcCompareResp, err error) {
	result = &qxTypesKms.KmsSkcCompareResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/compare", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}
