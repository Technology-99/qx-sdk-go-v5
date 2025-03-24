package qxKms

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxParser"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesCcs"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesKms"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	KmsBaseService interface {
		TestMsg(ctx context.Context, params *qxTypesKms.KmsTestMsgReq) (result *qxTypesKms.KmsTestMsgResp, err error)
	}

	defaultKmsBaseService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewKmsBaseService(qxCtx *qxCtx.QxCtx) KmsBaseService {
	// note: 初始化Kms系统
	return &defaultKmsBaseService{
		qxCtx: qxCtx,
	}
}

func (m *defaultKmsBaseService) TestMsg(ctx context.Context, params *qxTypesKms.KmsTestMsgReq) (result *qxTypesKms.KmsTestMsgResp, err error) {
	result = &qxTypesKms.KmsTestMsgResp{}
	if m.qxCtx.Parser.Status() != qxParser.QxParserStatusReady {
		logx.Errorf("ccs-TestMsg parser status not ready")
		return nil, nil
	}
	sendMsg, err := m.qxCtx.Parser.Encrypt(params.Msg)
	logx.Infof("打印加密后的消息: %s", sendMsg)
	if err != nil {
		logx.Errorf("ccs-TestMsg aes encrypt error: %v", err)
		return
	}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/testMsg", http.MethodPost, &qxTypesCcs.CcsTestMsgReq{
		Msg: sendMsg,
	})
	res, err := reqFn()
	if err != nil {
		logx.Errorf("ccs-TestMsg request error: %v", err)
		return nil, nil
	}

	tmpResultData := qxTypesKms.KmsEncryptResp{}
	_ = json.Unmarshal(res, &tmpResultData)
	if tmpResultData.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qiongxiao sdk errlog: ccs-TestMsg fail: %v", tmpResultData)
		result.Code = tmpResultData.Code
		result.Msg = qxCodes.StatusText(tmpResultData.Code)
		result.Path = tmpResultData.Path
		result.RequestID = tmpResultData.RequestID
		return result, nil
	}
	logx.Infof("ccs-TestMsg data: %v", tmpResultData.Data)
	// note: 使用aes解密数据
	decryptMsg, err := m.qxCtx.Parser.Decrypt(result.Msg)
	if err != nil {
		logx.Errorf("kms-TestMsg aes decrypt error: %v", err)
		return
	}
	logx.Infof("ccs-TestMsg decryptMsg: %v", string(decryptMsg))
	//aesResultData := qxTypesKms.KmsTestMsgRespData{}
	//if err = json.Unmarshal(decryptMsg, &aesResultData); err != nil {
	//	logx.Errorf("kms-TestMsg Data Unmarshal error: %v", err)
	//	return
	//}
	//result.Data = aesResultData

	return result, nil
}
