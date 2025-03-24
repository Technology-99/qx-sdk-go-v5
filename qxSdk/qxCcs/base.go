package qxCcs

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
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
		cli    *qxCli.QxClient
		parser qxParser.QxParser
	}
)

func NewCcsBaseService(cli *qxCli.QxClient, parser qxParser.QxParser) CcsBaseService {
	return &defaultCcsBaseService{
		cli:    cli,
		parser: parser,
	}
}

func (m *defaultCcsBaseService) TestMsg(ctx context.Context, params *qxTypesCcs.CcsTestMsgReq) (result *qxTypesCcs.CcsTestMsgResp, err error) {
	result = &qxTypesCcs.CcsTestMsgResp{}
	if m.parser.Status() != qxParser.QxParserStatusReady {
		logx.Errorf("ccs-TestMsg parser status not ready")
		return nil, nil
	}
	sendMsg, err := m.parser.Encrypt(params.Msg)
	logx.Infof("打印加密后的消息: %s", sendMsg)
	if err != nil {
		logx.Errorf("ccs-TestMsg aes encrypt error: %v", err)
		return
	}
	reqFn := m.cli.EasyNewRequest(ctx, "/ccs/testMsg", http.MethodPost, &qxTypesCcs.CcsTestMsgReq{
		Msg: sendMsg,
		Key: params.Key,
	})
	res, err := reqFn()
	if err != nil {
		logx.Errorf("ccs-TestMsg request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("qiongxiao sdk errlog: ccs-TestMsg fail: %v", result)
		return result, nil
	}
	// note: 使用aes解密数据
	aesResultData := qxTypesCcs.CcsTestMsgRespData{}
	decryptMsg, err := m.parser.Decrypt(result.Msg)
	if err != nil {
		logx.Errorf("ccs-TestMsg aes decrypt error: %v", err)
		return
	}
	if err = json.Unmarshal(decryptMsg, &aesResultData); err != nil {
		logx.Errorf("ccs-TestMsg Data Unmarshal error: %v", err)
		return
	}
	result.Data = aesResultData

	return result, nil
}
