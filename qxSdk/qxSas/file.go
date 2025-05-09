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
	FileService interface {
		// note: 文件管理部分
		// Create note: 创建一个文件
		Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error)
		// note: 添加一个文件并通过前端直传
		CreateAndDirectUpload(ctx context.Context, params *qxTypes.AllowCreateAndDirectUpload) (result *qxTypes.SasFileCreateAndDirectUploadResp, err error)
		// note: 上传完之后调用检查此接口完成上传
		CheckoutResult(ctx context.Context, params *qxTypes.SasFileApiCheckoutResultReq) (result *qxTypes.SasFileApiCheckoutResultResp, err error)
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

	defaultFileService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewFileService(qxCtx *qxCtx.QxCtx) FileService {
	return &defaultFileService{
		qxCtx: qxCtx,
	}
}

func (m *defaultFileService) Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error) {
	result = &qxTypes.SasFileApiCreateResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/create", http.MethodPost, &params)

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

func (m *defaultFileService) CreateAndDirectUpload(ctx context.Context, params *qxTypes.AllowCreateAndDirectUpload) (result *qxTypes.SasFileCreateAndDirectUploadResp, err error) {
	result = &qxTypes.SasFileCreateAndDirectUploadResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/createAndDirectUpload", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: CreateWithOssFrontUpload fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFileService) CheckoutResult(ctx context.Context, params *qxTypes.SasFileApiCheckoutResultReq) (result *qxTypes.SasFileApiCheckoutResultResp, err error) {
	result = &qxTypes.SasFileApiCheckoutResultResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/checkoutResult", http.MethodPost, &params)

	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: CreateWithOssFrontUpload fail: %v", result)
		return result, nil
	}
	return result, nil
}

func (m *defaultFileService) Delete(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/delete", http.MethodDelete, &params)

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

func (m *defaultFileService) DeleteMany(ctx context.Context, params *qxTypes.SasFileApiJsonIdsReq) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/deleteMany", http.MethodPost, &params)

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

func (m *defaultFileService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelSasFile) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/update", http.MethodPut, &params)

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

func (m *defaultFileService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelSasFile) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/updateStatus", http.MethodPatch, &params)

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

func (m *defaultFileService) Query(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileCommonQueryResp, err error) {
	result = &qxTypes.SasFileCommonQueryResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/query", http.MethodGet, &params)

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

func (m *defaultFileService) QueryListWhereIds(ctx context.Context, params *qxTypes.SasFileApiJsonIdsReq) (result *qxTypes.SasFileCommonQueryListResp, err error) {
	result = &qxTypes.SasFileCommonQueryListResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/queryListWhereIds", http.MethodPost, &params)

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

func (m *defaultFileService) QueryList(ctx context.Context, params *qxTypes.SasFileCommonSearchParams) (result *qxTypes.SasFileCommonQueryListResp, err error) {
	result = &qxTypes.SasFileCommonQueryListResp{}
	res, err := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/file/queryList", http.MethodPost, &params)

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
