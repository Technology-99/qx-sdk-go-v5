package qxBase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/google/go-querystring/query"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	QxBaseService interface {
		Codes(ctx context.Context, params *qxTypes.CodesReq) (result *qxTypes.CodesResp, err error)
		Zones(ctx context.Context, params *qxTypes.ZonesReq) (result *qxTypes.ZonesResp, err error)
	}
	defaultQxBaseService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewQxBaseService(qxCtx *qxCtx.QxCtx) QxBaseService {
	return &defaultQxBaseService{
		qxCtx: qxCtx,
	}
}

func (m *defaultQxBaseService) Codes(ctx context.Context, params *qxTypes.CodesReq) (result *qxTypes.CodesResp, err error) {
	result = &qxTypes.CodesResp{}
	relativePath := ""
	if params.Lang != "" || params.Svc != "" {
		v, _ := query.Values(params)
		relativePath = fmt.Sprintf("/codes?%s", v.Encode())
	} else {
		relativePath = "/codes"
	}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, relativePath, http.MethodGet, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: codes fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultQxBaseService) Zones(ctx context.Context, params *qxTypes.ZonesReq) (result *qxTypes.ZonesResp, err error) {
	result = &qxTypes.ZonesResp{}
	relativePath := ""
	if params.Lang != "" {
		v, _ := query.Values(params)
		relativePath = fmt.Sprintf("/codes?%s", v.Encode())
	} else {
		relativePath = "/codes"
	}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, relativePath, http.MethodGet, nil)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: zones fail: %v", result)
		return result, nil
	}
	return result, nil
}
