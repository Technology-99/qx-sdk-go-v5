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
	SasBaseService interface {
		QueryBucket(ctx context.Context, params *qxTypes.SasQueryBucketReq) (result *qxTypes.SasQueryBucketResp, err error)
	}
	defaultSasBaseService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewSasBaseService(qxCtx *qxCtx.QxCtx) SasBaseService {
	return &defaultSasBaseService{
		qxCtx: qxCtx,
	}
}

func (m *defaultSasBaseService) QueryBucket(ctx context.Context, params *qxTypes.SasQueryBucketReq) (result *qxTypes.SasQueryBucketResp, err error) {
	result = &qxTypes.SasQueryBucketResp{}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/sas/queryBucket", http.MethodPost, &params)
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
