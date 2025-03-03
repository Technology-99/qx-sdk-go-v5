package wechat

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	OffiaccountService interface {
		Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error)
	}

	defaultOffiaccountService struct {
		cli *qxCli.QxClient
	}
)

func NewOffiaccountService(cli *qxCli.QxClient) OffiaccountService {
	return &defaultOffiaccountService{
		cli: cli,
	}
}

func (m *defaultOffiaccountService) Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error) {
	result = &qxTypes.SasFileApiCreateResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/create", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("healthz request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != response.SUCCESS {
		logx.Errorf("qiongxiao sdk errlog: captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}
