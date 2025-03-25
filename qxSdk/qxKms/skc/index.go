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
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/createKeychain", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsSkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcEncryptReq) (result *qxTypesKms.KmsSkcEncryptResp, err error) {
	result = &qxTypesKms.KmsSkcEncryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/encrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsSkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcDecryptReq) (result *qxTypesKms.KmsSkcDecryptResp, err error) {
	result = &qxTypesKms.KmsSkcDecryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/decrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsSkcDecrypt error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsSkcDecrypt fail: %v", result)
		return result, nil
	}
	// note: 解密的数据需要额外解密
	decryptMsg, err := m.qxCtx.Parser.Decrypt(result.Data.BaseData)
	if err != nil {
		logx.Errorf("qx sdk KmsSkcDecrypt parser error: %v", err)
		return nil, nil
	}
	result.Data.BaseData = base64.StdEncoding.EncodeToString(decryptMsg)
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchEncrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchEncryptReq) (result *qxTypesKms.KmsSkcBatchEncryptResp, err error) {
	result = &qxTypesKms.KmsSkcBatchEncryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/batchEncrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsSkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultKmsSkcService) KmsSkcBatchDecrypt(ctx context.Context, params *qxTypesKms.KmsSkcBatchDecryptReq) (result *qxTypesKms.KmsSkcBatchDecryptResp, err error) {
	result = &qxTypesKms.KmsSkcBatchDecryptResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/batchDecrypt", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsSkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	for k, v := range result.Data {
		// note: 解密的数据需要额外解密
		decryptMsg, err := m.qxCtx.Parser.Decrypt(v.BaseData)
		if err != nil {
			logx.Errorf("qx sdk KmsSkcDecrypt parser error: %v", err)
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

func (m *defaultKmsSkcService) KmsSkcCompare(ctx context.Context, params *qxTypesKms.KmsSkcCompareReq) (result *qxTypesKms.KmsSkcCompareResp, err error) {
	result = &qxTypesKms.KmsSkcCompareResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/skc/compare", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk KmsSkcCreateKeychain error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk KmsSkcCreateKeychain fail: %v", result)
		return result, nil
	}
	for i, v := range result.Data.List {
		// note: 解密的数据需要额外解密
		decryptMsg, err := m.qxCtx.Parser.Decrypt(v.BaseData)
		if err != nil {
			logx.Errorf("qx sdk KmsSkcDecrypt parser error: %v", err)
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
