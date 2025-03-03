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
	CommonService interface {
		// note: 文件管理部分
		// Create note: 创建一个文件
		Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error)
		// CreateWithOssFrontUpload note: 添加一个文件并通过OSS前端直传
		CreateWithOssFrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelSasFileWithFrontedUpload) (result *qxTypes.SasFileCommonCreateWithOssFrontUploadResp, err error)
		// CreateWithOssV4FrontUpload note: 添加一个文件并通过OSSV4前端直传
		CreateWithOssV4FrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelSasFileWithFrontedUpload) (result *qxTypes.SasFileApiCreateWithOssV4FrontUploadResp, err error)
		// Delete note: 删除一个
		Delete(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileApiOKResp, err error)
		// DeleteMany note: 批量删除
		DeleteMany(ctx context.Context, params *qxTypes.SasFileApiFormIdsReq) (result *qxTypes.SasFileApiOKResp, err error)
		// Update note: 修改基础数据
		Update(ctx context.Context, params *qxTypes.AllowUpdateModelSasFile) (result *qxTypes.SasFileApiOKResp, err error)
		// UpdateStatus note: 修改启用状态
		UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelSasFile) (result *qxTypes.SasFileApiOKResp, err error)
		// Query note: 查询一个
		Query(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileCommonQueryResp, err error)
		// QueryListWhereIds note: 查询列表根据ids
		QueryListWhereIds(ctx context.Context, params *qxTypes.SasFileApiFormIdsReq) (result *qxTypes.SasFileCommonQueryListResp, err error)
		// QueryList note: 查询列表
		QueryList(ctx context.Context, params *qxTypes.SasFileCommonSearchParams) (result *qxTypes.SasFileCommonQueryListResp, err error)
	}

	defaultCommonService struct {
		cli *qxCli.QxClient
	}
)

func NewCommonService(cli *qxCli.QxClient) CommonService {
	return &defaultCommonService{
		cli: cli,
	}
}

func (m *defaultCommonService) Create(ctx context.Context, params *qxTypes.AllowCreateModelSasFile) (result *qxTypes.SasFileApiCreateResp, err error) {
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

func (m *defaultCommonService) CreateWithOssFrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelSasFileWithFrontedUpload) (result *qxTypes.SasFileCommonCreateWithOssFrontUploadResp, err error) {
	result = &qxTypes.SasFileCommonCreateWithOssFrontUploadResp{}
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

func (m *defaultCommonService) CreateWithOssV4FrontUpload(ctx context.Context, params *qxTypes.AllowCreateModelSasFileWithFrontedUpload) (result *qxTypes.SasFileApiCreateWithOssV4FrontUploadResp, err error) {
	result = &qxTypes.SasFileApiCreateWithOssV4FrontUploadResp{}
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

func (m *defaultCommonService) Delete(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
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

func (m *defaultCommonService) DeleteMany(ctx context.Context, params *qxTypes.SasFileApiFormIdsReq) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
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

func (m *defaultCommonService) Update(ctx context.Context, params *qxTypes.AllowUpdateModelSasFile) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
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

func (m *defaultCommonService) UpdateStatus(ctx context.Context, params *qxTypes.AllowUpdateStatusModelSasFile) (result *qxTypes.SasFileApiOKResp, err error) {
	result = &qxTypes.SasFileApiOKResp{}
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

func (m *defaultCommonService) Query(ctx context.Context, params *qxTypes.SasFileApiFormIdReq) (result *qxTypes.SasFileCommonQueryResp, err error) {
	result = &qxTypes.SasFileCommonQueryResp{}
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

func (m *defaultCommonService) QueryListWhereIds(ctx context.Context, params *qxTypes.SasFileApiFormIdsReq) (result *qxTypes.SasFileCommonQueryListResp, err error) {
	result = &qxTypes.SasFileCommonQueryListResp{}
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

func (m *defaultCommonService) QueryList(ctx context.Context, params *qxTypes.SasFileCommonSearchParams) (result *qxTypes.SasFileCommonQueryListResp, err error) {
	result = &qxTypes.SasFileCommonQueryListResp{}
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
