package qxCcs

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxParser"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesCcs"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	CcsBaseService interface {
		TestMsg(ctx context.Context, params *qxTypesCcs.CcsTestMsgReq) (result *qxTypesCcs.CcsTestMsgResp, err error)
	}
	defaultCcsBaseService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewCcsBaseService(qxCtx *qxCtx.QxCtx) CcsBaseService {
	return &defaultCcsBaseService{
		qxCtx: qxCtx,
	}
}

func (m *defaultCcsBaseService) TestMsg(ctx context.Context, params *qxTypesCcs.CcsTestMsgReq) (result *qxTypesCcs.CcsTestMsgResp, err error) {
	result = &qxTypesCcs.CcsTestMsgResp{}
	if m.qxCtx.Parser.Status() != qxParser.QxParserStatusReady {
		logx.Errorf("qx sdk: parser status not ready")
		return nil, nil
	}
	sendMsg, err := m.qxCtx.Parser.Encrypt(params.Msg)
	logx.Infof("打印加密后的消息: %s", sendMsg)
	if err != nil {
		logx.Errorf("qx sdk: aes encrypt error: %v", err)
		return
	}
	reqFn := m.qxCtx.Cli.EasyNewRequest(ctx, "/ccs/testMsg", http.MethodPost, &qxTypesCcs.CcsTestMsgReq{
		Msg: sendMsg,
		Key: params.Key,
	})
	res, err := reqFn()
	if err != nil {
		logx.Errorf("qx sdk: request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qx sdk: ccs-TestMsg fail: %v", result)
		return result, nil
	}
	// note: 使用aes解密数据
	aesResultData := qxTypesCcs.CcsTestMsgRespData{}
	decryptMsg, err := m.qxCtx.Parser.Decrypt(result.Msg)
	if err != nil {
		logx.Errorf("qx sdk: decrypt error: %v", err)
		return
	}
	if err = json.Unmarshal(decryptMsg, &aesResultData); err != nil {
		logx.Errorf("qx sdk: Unmarshal error: %v", err)
		return
	}
	result.Data = aesResultData

	return result, nil
}
