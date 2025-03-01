package qxSas

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
	FileCollService interface {
		// note: 文件管理部分
		// Create note: 创建一个文件
		Create(ctx context.Context, params *qxTypes.AllowCreateModelTmsFile) (result *qxTypes.TmsFileApiCreateResp, err error)
	}

	defaultFileCollService struct {
		cli *qxCli.QxClient
	}
)

func NewFileCollService(cli *qxCli.QxClient) FileCollService {
	return &defaultFileCollService{
		cli: cli,
	}
}

func (m *defaultFileCollService) Create(ctx context.Context, params *qxTypes.AllowCreateModelTmsFile) (result *qxTypes.TmsFileApiCreateResp, err error) {
	result = &qxTypes.TmsFileApiCreateResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/fileColl/create", http.MethodPost, &params)
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
