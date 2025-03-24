package qxParser

import (
	"context"
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qxLib/qxCodes"
	"github.com/Technology-99/qxLib/qxCrypto"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/hkdf"
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
	m.status = QxParserStatusInit
	m.mu.Lock()
	defer m.mu.Unlock()

	privKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		logx.Errorf("无法生成公钥: %s", err)
		m.status = QxParserStatusNoInit
		return err
	}
	m.localPrivateKey = privKey
	ctx := context.Background()

	tmpPub := base64.StdEncoding.EncodeToString(m.localPrivateKey.PublicKey().Bytes())
	// 发送公钥给网关
	reqFn := m.cli.EasyNewRequest(ctx, "/keyExchange", http.MethodPost, &qxTypes.QxClientKeyExChangeReq{
		AccessKey: m.cli.AccessKeyId,
		PublicKey: tmpPub,
	})
	res, err := reqFn()
	if err != nil {
		logx.Errorf("keyexchange request error: %v", err)
		m.status = QxParserStatusNoInit
		return err
	}
	result := qxTypes.QxClientKeyExChangeResp{}
	_ = json.Unmarshal(res, &result)
	if result.Code != qxCodes.QxEngineStatusOK {
		logx.Errorf("keyexchange  fail: %v", result)
		m.status = QxParserStatusNoInit
		return errors.New(result.Msg)
	}
	logx.Infof("keyexchange request success: %v", result)
	// note: 完成解析器的初始化

	tmpPubkey, err := base64.StdEncoding.DecodeString(result.Data.PublicKey)
	if err != nil {
		logx.Errorf("keyexchange load tmpPubkey error: %v", err)
		m.status = QxParserStatusNoInit
		return err
	}

	// 解析远端公钥
	remotePubKey, err := ecdh.P256().NewPublicKey([]byte(tmpPubkey))
	if err != nil {
		logx.Errorf("keyexchange parse remote public key error: %v", err)
		m.status = QxParserStatusNoInit
		return err
	}

	// 计算共享密钥
	sharedSecret, err := m.localPrivateKey.ECDH(remotePubKey)
	if err != nil {
		logx.Errorf("keyexchange compute share key error: %v", err)
		m.status = QxParserStatusNoInit
		return err
	}

	// 通过 HKDF 派生 AES-256 密钥
	hkdf := hkdf.New(sha256.New, sharedSecret, nil, []byte("key-session-key"))
	derivedKey := make([]byte, qxCrypto.AES256KeyLen) // 32 字节 AES-256 密钥
	hkdf.Read(derivedKey)

	// 更新 KMS 客户端状态
	m.sharedSecretBase = base64.StdEncoding.EncodeToString(sharedSecret)
	m.remotePublicKey = remotePubKey
	m.sharedSecretBase = base64.StdEncoding.EncodeToString(sharedSecret)
	m.deriveAESKeyBase = base64.StdEncoding.EncodeToString(derivedKey)
	m.deriveAESIvBase = base64.StdEncoding.EncodeToString(sharedSecret[0:qxCrypto.AESGCMIvLen])

	m.status = QxParserStatusReady

	logx.Infof("keyexchange success")
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
