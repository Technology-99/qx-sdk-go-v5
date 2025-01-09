package media

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/cli"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/types"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	FileService interface {
		// Create note: 创建一个文件
		Create(ctx context.Context, params *types.AllowCreateModelTmsFile) (result *types.TmsFileApiCreateResp, err error)
		// CreateWithOssFrontUpload note: 添加一个文件并通过OSS前端直传
		CreateWithOssFrontUpload(ctx context.Context, params *types.AllowCreateModelTmsFileWithFrontedUpload) (result *types.TmsFileCommonCreateWithOssFrontUploadResp, err error)
		// CreateWithOssV4FrontUpload note: 添加一个文件并通过OSSV4前端直传
		CreateWithOssV4FrontUpload(ctx context.Context, params *types.AllowCreateModelTmsFileWithFrontedUpload) (result *types.TmsFileApiCreateWithOssV4FrontUploadResp, err error)
		// Delete note: 删除一个
		Delete(ctx context.Context, params *types.TmsFileApiFormIdReq) (result *types.TmsFileApiOKResp, err error)
		// DeleteMany note: 批量删除
		DeleteMany(ctx context.Context, params *types.TmsFileApiFormIdsReq) (result *types.TmsFileApiOKResp, err error)
		// Update note: 修改基础数据
		Update(ctx context.Context, params *types.AllowUpdateModelTmsFile) (result *types.TmsFileApiOKResp, err error)
		// UpdateStatus note: 修改启用状态
		UpdateStatus(ctx context.Context, params *types.AllowUpdateStatusModelTmsFile) (result *types.TmsFileApiOKResp, err error)
		// Query note: 查询一个
		Query(ctx context.Context, params *types.TmsFileApiFormIdReq) (result *types.TmsFileCommonQueryResp, err error)
		// QueryListWhereIds note: 查询列表根据ids
		QueryListWhereIds(ctx context.Context, params *types.TmsFileApiFormIdsReq) (result *types.TmsFileCommonQueryListResp, err error)
		// QueryList note: 查询列表
		QueryList(ctx context.Context, params *types.TmsFileCommonSearchParams) (result *types.TmsFileCommonQueryListResp, err error)
	}

	defaultFileService struct {
		cli *cli.QxClient
	}
)

func NewFileService(cli *cli.QxClient) FileService {
	return &defaultFileService{
		cli: cli,
	}
}

func (m *defaultFileService) Create(ctx context.Context, params *types.AllowCreateModelTmsFile) (result *types.TmsFileApiCreateResp, err error) {
	result = &types.TmsFileApiCreateResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/create", http.MethodPost, &params)
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

func (m *defaultFileService) CreateWithOssFrontUpload(ctx context.Context, params *types.AllowCreateModelTmsFileWithFrontedUpload) (result *types.TmsFileCommonCreateWithOssFrontUploadResp, err error) {
	result = &types.TmsFileCommonCreateWithOssFrontUploadResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/createWithOssFrontedUpload", http.MethodPost, &params)
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

func (m *defaultFileService) CreateWithOssV4FrontUpload(ctx context.Context, params *types.AllowCreateModelTmsFileWithFrontedUpload) (result *types.TmsFileApiCreateWithOssV4FrontUploadResp, err error) {
	result = &types.TmsFileApiCreateWithOssV4FrontUploadResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/createWithOssV4FrontedUpload", http.MethodPost, &params)
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

func (m *defaultFileService) Delete(ctx context.Context, params *types.TmsFileApiFormIdReq) (result *types.TmsFileApiOKResp, err error) {
	result = &types.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/delete", http.MethodDelete, &params)
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

func (m *defaultFileService) DeleteMany(ctx context.Context, params *types.TmsFileApiFormIdsReq) (result *types.TmsFileApiOKResp, err error) {
	result = &types.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/deleteMany", http.MethodDelete, &params)
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

func (m *defaultFileService) Update(ctx context.Context, params *types.AllowUpdateModelTmsFile) (result *types.TmsFileApiOKResp, err error) {
	result = &types.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/update", http.MethodPut, &params)
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

func (m *defaultFileService) UpdateStatus(ctx context.Context, params *types.AllowUpdateStatusModelTmsFile) (result *types.TmsFileApiOKResp, err error) {
	result = &types.TmsFileApiOKResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/updateStatus", http.MethodPatch, &params)
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

func (m *defaultFileService) Query(ctx context.Context, params *types.TmsFileApiFormIdReq) (result *types.TmsFileCommonQueryResp, err error) {
	result = &types.TmsFileCommonQueryResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/query", http.MethodGet, &params)
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

func (m *defaultFileService) QueryListWhereIds(ctx context.Context, params *types.TmsFileApiFormIdsReq) (result *types.TmsFileCommonQueryListResp, err error) {
	result = &types.TmsFileCommonQueryListResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/queryListWhereIds", http.MethodGet, &params)
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

func (m *defaultFileService) QueryList(ctx context.Context, params *types.TmsFileCommonSearchParams) (result *types.TmsFileCommonQueryListResp, err error) {
	result = &types.TmsFileCommonQueryListResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/tmsFile/queryList", http.MethodPost, &params)
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
