package qxParser

import (
	"context"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxCrypto"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
)

const (
	QxParserStatusNoInit = iota + 1
	QxParserStatusInit
	QxParserStatusReady
)

type (
	QxParser interface {
		Status() int
		Decrypt(data string) ([]byte, error)
		Encrypt(data string) (string, error)
		KeyExchange() error
	}
	defaultQxParser struct {
		cli              *qxCli.QxClient
		status           int
		mu               sync.Mutex
		localPrivateKey  *ecdh.PrivateKey
		remotePublicKey  *ecdh.PublicKey
		sharedSecretBase string
		deriveAESKeyBase string
		deriveAESIvBase  string
	}
)

func NewQxParser(cli *qxCli.QxClient) QxParser {
	return &defaultQxParser{
		cli:    cli,
		status: QxParserStatusNoInit,
	}
}

func (m *defaultQxParser) KeyExchange() error {
	privKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		logx.Errorf("无法生成公钥: %s", err)
		return err
	}
	m.localPrivateKey = privKey
	m.status = QxParserStatusInit
	ctx := context.Background()

	tmpPub := base64.StdEncoding.EncodeToString(m.localPrivateKey.PublicKey().Bytes())
	logx.Infof("打印一下公钥: %s", tmpPub)
	// 发送公钥给网关
	reqFn := m.cli.EasyNewRequest(ctx, "/keyExchange", http.MethodPost, &qxTypes.QxClientKeyExChangeReq{
		PublicKey: tmpPub,
	})
	res, err := reqFn()
	if err != nil {
		logx.Errorf("keyexchange request error: %v", err)
		return err
	}
	result := qxTypes.QxClientKeyExChangeResp{}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("keyexchange  fail: %v", result)
		return errors.New(result.Msg)
	}
	logx.Infof("keyexchange success: %v", result)
	return nil
}

func (m *defaultQxParser) Encrypt(data string) (string, error) {
	baseData, err := qxCrypto.AESEncryptByGCM([]byte(data), m.deriveAESKeyBase, m.deriveAESIvBase)
	return baseData, err
}

func (m *defaultQxParser) Decrypt(data string) ([]byte, error) {
	decodeData, err := qxCrypto.AESDecryptByGCM(data, m.deriveAESKeyBase, m.deriveAESIvBase)
	return decodeData, err
}

func (m *defaultQxParser) Status() int {
	return m.status
}
