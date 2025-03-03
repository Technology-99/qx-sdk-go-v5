package qxCms

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesCms"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxUtils"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	CmsBaseService interface {
		TestMsg(ctx context.Context, params *qxTypesCms.CmsTestMsgReq) (result *qxTypesCms.CmsTestMsgResp, err error)
	}
	defaultCmsBaseService struct {
		cli *qxCli.QxClient
	}
)

func NewCmsBaseService(cli *qxCli.QxClient) CmsBaseService {
	return &defaultCmsBaseService{
		cli: cli,
	}
}

func (m *defaultCmsBaseService) TestMsg(ctx context.Context, params *qxTypesCms.CmsTestMsgReq) (result *qxTypesCms.CmsTestMsgResp, err error) {
	result = &qxTypesCms.CmsTestMsgResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/cms/testMsg", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("CmsBaseService TestMsg request error: %v", err)
		return nil, err
	}
	logx.Infof("CmsBaseService TestMsg request response: %s", res)
	tmpResult := &qxTypesCms.CmsEncryptResp{}
	_ = json.Unmarshal(res, &tmpResult)
	if m.cli.Config.EncryptionPrivateKey != "" {
		// note: 先进行解密，再解析
		tmpDataBt, err := qxUtils.RsaDecrypt(tmpResult.Data, m.cli.Config.EncryptionPrivateKey)
		if err != nil {
			logx.Errorf("CmsBaseService TestMsg RsaDecrypt error: %v", err)
			return nil, err
		}
		if err = json.Unmarshal(tmpDataBt, &result.Data); err != nil {
			logx.Errorf("CmsBaseService TestMsg json.Unmarshal error: %v", err)
			return nil, err
		}
	} else {
		logx.Infof("Please contact your system administrator to download and deploy the decryption certificate")
		logx.Infof("CmsBaseService TestMsg response decrypt need encryptionPrivateKey")
	}

	result.Msg = tmpResult.Msg
	result.Code = tmpResult.Code
	result.Path = tmpResult.Path
	result.RequestID = tmpResult.RequestID
	if result.Code != response.SUCCESS {
		logx.Errorf("qiongxiao sdk errlog: TestMsg fail: %v", result)
		return result, nil
	}

	return result, nil
}
