package qxSas

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
	FolderService interface {
		// note: 文件管理部分
		// Create note: 创建一个文件夹
		Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFolder) (result *qxTypes.SasFolderApiCreateResp, err error)
		// Delete note: 删除一个
		Delete(ctx context.Context, params *qxTypes.SasFolderApiFormIdReq) (result *qxTypes.SasFolderApiOKResp, err error)
		// DeleteMany note: 批量删除
		DeleteMany(ctx context.Context, params *qxTypes.SasFolderApiJsonIdsReq) (result *qxTypes.SasFolderApiOKResp, err error)
		// Update note: 修改基础数据
		Update(ctx context.Context, params *qxTypes.AllowUpdateModelSasFolder) (result *qxTypes.SasFolderApiOKResp, err error)
		// UpdateStatus note: 修改启用状态
		UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelSasFolder) (result *qxTypes.SasFolderApiOKResp, err error)
		// Query note: 查询一个
		Query(ctx context.Context, params *qxTypes.SasFolderApiFormIdReq) (result *qxTypes.SasFolderCommonQueryResp, err error)
		// QueryListWhereIds note: 查询列表根据ids
		QueryListWhereIds(ctx context.Context, params *qxTypes.SasFolderApiJsonIdsReq) (result *qxTypes.SasFolderCommonQueryListResp, err error)
		// QueryList note: 查询列表
		QueryList(ctx context.Context, params *qxTypes.SasFolderCommonSearchParams) (result *qxTypes.SasFolderCommonQueryListResp, err error)
	}

	defaultFolderService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewFolderService(qxCtx *qxCtx.QxCtx) FolderService {
	return &defaultFolderService{
		qxCtx: qxCtx,
	}
}

func (m *defaultFolderService) Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFolder) (result *qxTypes.SasFolderApiCreateResp, err error) {
	result = &qxTypes.SasFolderApiCreateResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/create", http.MethodPost, &params)

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

func (m *defaultFolderService) Delete(ctx context.Context, params *qxTypes.SasFolderApiFormIdReq) (result *qxTypes.SasFolderApiOKResp, err error) {
	result = &qxTypes.SasFolderApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/delete", http.MethodDelete, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: Delete fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFolderService) DeleteMany(ctx context.Context, params *qxTypes.SasFolderApiJsonIdsReq) (result *qxTypes.SasFolderApiOKResp, err error) {
	result = &qxTypes.SasFolderApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/deleteMany", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: DeleteMany fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFolderService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelSasFolder) (result *qxTypes.SasFolderApiOKResp, err error) {
	result = &qxTypes.SasFolderApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/update", http.MethodPut, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: Update fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFolderService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelSasFolder) (result *qxTypes.SasFolderApiOKResp, err error) {
	result = &qxTypes.SasFolderApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/updateStatus", http.MethodPatch, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: UpdateStatus fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFolderService) Query(ctx context.Context, params *qxTypes.SasFolderApiFormIdReq) (result *qxTypes.SasFolderCommonQueryResp, err error) {
	result = &qxTypes.SasFolderCommonQueryResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/query", http.MethodGet, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: Query fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFolderService) QueryListWhereIds(ctx context.Context, params *qxTypes.SasFolderApiJsonIdsReq) (result *qxTypes.SasFolderCommonQueryListResp, err error) {
	result = &qxTypes.SasFolderCommonQueryListResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/queryListWhereIds", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: QueryListWhereIds fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFolderService) QueryList(ctx context.Context, params *qxTypes.SasFolderCommonSearchParams) (result *qxTypes.SasFolderCommonQueryListResp, err error) {
	result = &qxTypes.SasFolderCommonQueryListResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/folder/queryList", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: QueryList fail: %v", result)
		return result, nil
	}
	return result, nil
}
