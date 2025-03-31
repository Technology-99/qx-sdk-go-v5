package wechat

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
	OffiaccountService interface {
		Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error)
	}

	defaultOffiaccountService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewOffiaccountService(qxCtx *qxCtx.QxCtx) OffiaccountService {
	return &defaultOffiaccountService{
		qxCtx: qxCtx,
	}
}

func (m *defaultOffiaccountService) Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error) {
	result = &qxTypes.SasFileApiCreateResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/create", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: Create fail: %v", result)
		return result, nil
	}
	return result, nil
}
