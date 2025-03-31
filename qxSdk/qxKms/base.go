package qxKms

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxParser"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesCcs"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesKms"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
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
		logx.Errorf("qx sdk: parser status not ready")
		return nil, nil
	}
	msg := fmt.Sprintf("test msg: %s", time.Now().Format(time.DateTime))
	encryptMsg, err := m.qxCtx.Parser.Encrypt(msg)
	if err != nil {
		logx.Errorf("qx sdk: aes encrypt error: %v", err)
		return
	}
	if err != nil {
		logx.Errorf("qx sdk: aes encrypt error: %v", err)
		return
	}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/kms/testMsg", http.MethodPost, &qxTypesCcs.CcsTestMsgReq{
		Msg: encryptMsg,
	})
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	tmpResultData := qxTypesKms.KmsEncryptResp{}
	_ = json.Unmarshal(res, &tmpResultData)
	if tmpResultData.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: kms-TestMsg fail: %v", tmpResultData)
		result.Code = tmpResultData.Code
		result.Msg = qxCodes.StatusText(tmpResultData.Code)
		result.Path = tmpResultData.Path
		result.RequestID = tmpResultData.RequestID
		return result, nil
	}
	logx.Infof("打印解析器的状态 %s", m.qxCtx.Parser.FormatStatus())
	// note: 使用aes解密数据
	decryptMsg, err := m.qxCtx.Parser.Decrypt(tmpResultData.Data)
	if err != nil {
		logx.Errorf("qx sdk: aes decrypt error: %v", err)
		return
	}
	logx.Infof("qx sdk: decryptMsg: %v", string(decryptMsg))
	aesResultData := qxTypesKms.KmsTestMsgRespData{}
	if err = json.Unmarshal(decryptMsg, &aesResultData); err != nil {
		logx.Errorf("qx sdk: Data Unmarshal error: %v", err)
		return
	}
	result.Data = aesResultData

	return result, nil
}
