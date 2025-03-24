package aes

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesCcs"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	AesService interface {
		// note: 根据key获取aes密钥
		GetAesFromKey(ctx context.Context, params *qxTypesCcs.GetAesFromKeyReq) (result *qxTypesCcs.GetAesFromKeyResp, err error)
		// note: 获取最新的aes密钥
		GetAesLatest(ctx context.Context, params *qxTypesCcs.GetAesLatestReq) (result *qxTypesCcs.GetAesLatestResp, err error)
	}

	defaultAesService struct {
		cli *qxCli.QxClient
	}
)

func NewAesService(cli *qxCli.QxClient) AesService {
	return &defaultAesService{
		cli: cli,
	}
}

func (m *defaultAesService) GetAesFromKey(ctx context.Context, params *qxTypesCcs.GetAesFromKeyReq) (result *qxTypesCcs.GetAesFromKeyResp, err error) {
	result = &qxTypesCcs.GetAesFromKeyResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/ccs/aes/getAesFromKey", http.MethodPost, &params)
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

func (m *defaultAesService) GetAesLatest(ctx context.Context, params *qxTypesCcs.GetAesLatestReq) (result *qxTypesCcs.GetAesLatestResp, err error) {
	result = &qxTypesCcs.GetAesLatestResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/ccs/aes/getAesLatest", http.MethodPost, &params)
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
