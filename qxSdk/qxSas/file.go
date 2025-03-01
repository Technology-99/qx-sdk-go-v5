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
	FileService interface {
		// note: 文件管理部分
		// Create note: 创建一个文件
		Create(ctx context.Context, params *qxTypes.AllowCreateModelTmsFile) (result *qxTypes.TmsFileApiCreateResp, err error)
		// CreateWithOssFrontUpload note: 添加一个文件并通过OSS前端直传
		CreateWithOssFrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelTmsFileWithFrontedUpload) (result *qxTypes.TmsFileCommonCreateWithOssFrontUploadResp, err error)
		// CreateWithOssV4FrontUpload note: 添加一个文件并通过OSSV4前端直传
		CreateWithOssV4FrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelTmsFileWithFrontedUpload) (result *qxTypes.TmsFileApiCreateWithOssV4FrontUploadResp, err error)
		// Delete note: 删除一个
		Delete(ctx context.Context, params *qxTypes.TmsFileApiFormIdReq) (result *qxTypes.TmsFileApiOKResp, err error)
		// DeleteMany note: 批量删除
		DeleteMany(ctx context.Context, params *qxTypes.TmsFileApiFormIdsReq) (result *qxTypes.TmsFileApiOKResp, err error)
		// Update note: 修改基础数据
		Update(ctx context.Context, params *qxTypes.AllowUpdateModelTmsFile) (result *qxTypes.TmsFileApiOKResp, err error)
		// UpdateStatus note: 修改启用状态
		UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelTmsFile) (result *qxTypes.TmsFileApiOKResp, err error)
		// Query note: 查询一个
		Query(ctx context.Context, params *qxTypes.TmsFileApiFormIdReq) (result *qxTypes.TmsFileCommonQueryResp, err error)
		// QueryListWhereIds note: 查询列表根据ids
		QueryListWhereIds(ctx context.Context, params *qxTypes.TmsFileApiFormIdsReq) (result *qxTypes.TmsFileCommonQueryListResp, err error)
		// QueryList note: 查询列表
		QueryList(ctx context.Context, params *qxTypes.TmsFileCommonSearchParams) (result *qxTypes.TmsFileCommonQueryListResp, err error)
	}

	defaultFileService struct {
		cli *qxCli.QxClient
	}
)

func NewFileService(cli *qxCli.QxClient) FileService {
	return &defaultFileService{
		cli: cli,
	}
}

func (m *defaultFileService) Create(ctx context.Context, params *qxTypes.AllowCreateModelTmsFile) (result *qxTypes.TmsFileApiCreateResp, err error) {
	result = &qxTypes.TmsFileApiCreateResp{}
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

func (m *defaultFileService) CreateWithOssFrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelTmsFileWithFrontedUpload) (result *qxTypes.TmsFileCommonCreateWithOssFrontUploadResp, err error) {
	result = &qxTypes.TmsFileCommonCreateWithOssFrontUploadResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/createWithOssFrontedUpload", http.MethodPost, &params)
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

func (m *defaultFileService) CreateWithOssV4FrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelTmsFileWithFrontedUpload) (result *qxTypes.TmsFileApiCreateWithOssV4FrontUploadResp, err error) {
	result = &qxTypes.TmsFileApiCreateWithOssV4FrontUploadResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/createWithOssV4FrontedUpload", http.MethodPost, &params)
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

func (m *defaultFileService) Delete(ctx context.Context, params *qxTypes.TmsFileApiFormIdReq) (result *qxTypes.TmsFileApiOKResp, err error) {
	result = &qxTypes.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/delete", http.MethodDelete, &params)
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

func (m *defaultFileService) DeleteMany(ctx context.Context, params *qxTypes.TmsFileApiFormIdsReq) (result *qxTypes.TmsFileApiOKResp, err error) {
	result = &qxTypes.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/deleteMany", http.MethodDelete, &params)
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

func (m *defaultFileService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelTmsFile) (result *qxTypes.TmsFileApiOKResp, err error) {
	result = &qxTypes.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/update", http.MethodPut, &params)
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

func (m *defaultFileService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelTmsFile) (result *qxTypes.TmsFileApiOKResp, err error) {
	result = &qxTypes.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/updateStatus", http.MethodPatch, &params)
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

func (m *defaultFileService) Query(ctx context.Context, params *qxTypes.TmsFileApiFormIdReq) (result *qxTypes.TmsFileCommonQueryResp, err error) {
	result = &qxTypes.TmsFileCommonQueryResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/query", http.MethodGet, &params)
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

func (m *defaultFileService) QueryListWhereIds(ctx context.Context, params *qxTypes.TmsFileApiFormIdsReq) (result *qxTypes.TmsFileCommonQueryListResp, err error) {
	result = &qxTypes.TmsFileCommonQueryListResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/queryListWhereIds", http.MethodGet, &params)
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

func (m *defaultFileService) QueryList(ctx context.Context, params *qxTypes.TmsFileCommonSearchParams) (result *qxTypes.TmsFileCommonQueryListResp, err error) {
	result = &qxTypes.TmsFileCommonQueryListResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/sas/file/queryList", http.MethodPost, &params)
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
