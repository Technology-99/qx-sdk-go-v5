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
	CommonService interface {
		// note: 文件管理部分
		// Create note: 创建一个文件
		Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error)
		// Delete note: 删除一个
		Delete(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileApiOKResp, err error)
		// DeleteMany note: 批量删除
		DeleteMany(ctx context.Context, params *qxTypes.SasFileApiJsonIdsReq) (result *qxTypes.SasFileApiOKResp, err error)
		// Update note: 修改基础数据
		Update(ctx context.Context, params *qxTypes.AllowUpdateModelSasFile) (result *qxTypes.SasFileApiOKResp, err error)
		// UpdateStatus note: 修改启用状态
		UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelSasFile) (result *qxTypes.SasFileApiOKResp, err error)
		// Query note: 查询一个
		Query(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileCommonQueryResp, err error)
		// QueryListWhereIds note: 查询列表根据ids
		QueryListWhereIds(ctx context.Context, params *qxTypes.SasFileApiJsonIdsReq) (result *qxTypes.SasFileCommonQueryListResp, err error)
		// QueryList note: 查询列表
		QueryList(ctx context.Context, params *qxTypes.SasFileCommonSearchParams) (result *qxTypes.SasFileCommonQueryListResp, err error)
	}

	defaultCommonService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewCommonService(qxCtx *qxCtx.QxCtx) CommonService {
	return &defaultCommonService{
		qxCtx: qxCtx,
	}
}

func (m *defaultCommonService) Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error) {
	result = &qxTypes.SasFileApiCreateResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/create", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultCommonService) Delete(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/delete", http.MethodDelete, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultCommonService) DeleteMany(ctx context.Context, params *qxTypes.SasFileApiJsonIdsReq) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/deleteMany", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultCommonService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelSasFile) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/update", http.MethodPut, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultCommonService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelSasFile) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/updateStatus", http.MethodPatch, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultCommonService) Query(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileCommonQueryResp, err error) {
	result = &qxTypes.SasFileCommonQueryResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/query", http.MethodGet, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultCommonService) QueryListWhereIds(ctx context.Context, params *qxTypes.SasFileApiJsonIdsReq) (result *qxTypes.SasFileCommonQueryListResp, err error) {
	result = &qxTypes.SasFileCommonQueryListResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/queryListWhereIds", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultCommonService) QueryList(ctx context.Context, params *qxTypes.SasFileCommonSearchParams) (result *qxTypes.SasFileCommonQueryListResp, err error) {
	result = &qxTypes.SasFileCommonQueryListResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/queryList", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk:captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}
