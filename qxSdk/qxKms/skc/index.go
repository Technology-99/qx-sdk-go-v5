package skc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesKms"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	KmsSkcService interface {
		KmsAkcCreateKeychain(ctx context.Context, params *qxTypesKms.KmsSkcCreateKeychainReq) (result *qxTypesKms.KmsSkcCreateKeychainResp, err error)
		KmsAkcEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcEncryptReq) (result *qxTypesKms.KmsSkcEncryptResp, err error)
		KmsAkcDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcDecryptReq) (result *qxTypesKms.KmsSkcDecryptResp, err error)
		KmsAkcBatchEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchEncryptReq) (result *qxTypesKms.KmsSkcBatchEncryptResp, err error)
		KmsAkcBatchDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchDecryptReq) (result *qxTypesKms.KmsSkcBatchDecryptResp, err error)
		KmsAkcCompare(ctx context.Context, params *qxTypesKms.KmsSkcCompareReq) (result *qxTypesKms.KmsSkcCompareResp, err error)
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

func (m *defaultKmsSkcService) KmsAkcCreateKeychain(ctx context.Context, params *qxTypesKms.KmsSkcCreateKeychainReq) (result *qxTypesKms.KmsSkcCreateKeychainResp, err error) {
	result = &qxTypesKms.KmsSkcCreateKeychainResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/createKeychain", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsAkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsAkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsAkcEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcEncryptReq) (result *qxTypesKms.KmsSkcEncryptResp, err error) {
	result = &qxTypesKms.KmsSkcEncryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/encrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsAkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsAkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsAkcDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcDecryptReq) (result *qxTypesKms.KmsSkcDecryptResp, err error) {
	result = &qxTypesKms.KmsSkcDecryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/decrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsAkcDecrypt error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsAkcDecrypt fail: %v", result)
		return result, nil
	}
	// note: 解密的数据需要额外解密
	decryptMsg, err := m.qxCtx.Parser.Decrypt(result.Data.BaseData)
	if err != nil {
		logx.Errorf("qx sdk KmsAkcDecrypt parser error: %v", err)
		return nil, nil
	}
	result.Data.BaseData = base64.StdEncoding.EncodeToString(decryptMsg)
	return result, nil
}

func (m *defaultKmsSkcService) KmsAkcBatchEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchEncryptReq) (result *qxTypesKms.KmsSkcBatchEncryptResp, err error) {
	result = &qxTypesKms.KmsSkcBatchEncryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/batchEncrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsAkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsAkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsAkcBatchDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchDecryptReq) (result *qxTypesKms.KmsSkcBatchDecryptResp, err error) {
	result = &qxTypesKms.KmsSkcBatchDecryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/batchDecrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsAkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsAkcCreateKeychain fail: %v", result)
		return result, nil
	}
	for k, v := range result.Data {
		// note: 解密的数据需要额外解密
		decryptMsg, err := m.qxCtx.Parser.Decrypt(v.BaseData)
		if err != nil {
			logx.Errorf("qx sdk KmsAkcDecrypt parser error: %v", err)
			return nil, nil
		}
		result.Data[k] = qxTypesKms.KmsSkcBatchDecryptRespItem{
			Name:     v.Name,
			BaseData: base64.StdEncoding.EncodeToString(decryptMsg),
			Status:   v.Status,
		}
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsAkcCompare(ctx context.Context, params *qxTypesKms.KmsSkcCompareReq) (result *qxTypesKms.KmsSkcCompareResp, err error) {
	result = &qxTypesKms.KmsSkcCompareResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/compare", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsAkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsAkcCreateKeychain fail: %v", result)
		return result, nil
	}
	for i, v := range result.Data.List {
		// note: 解密的数据需要额外解密
		decryptMsg, err := m.qxCtx.Parser.Decrypt(v.BaseData)
		if err != nil {
			logx.Errorf("qx sdk KmsAkcDecrypt parser error: %v", err)
			return nil, nil
		}
		result.Data.List[i] = qxTypesKms.KmsSkcCompareRespDataItem{
			Name:     v.Name,
			BaseData: base64.StdEncoding.EncodeToString(decryptMsg),
			Status:   v.Status,
		}
	}
	return result, nil
}
