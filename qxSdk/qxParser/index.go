package qxParser

import (
	"crypto/ecdh"
	"crypto/sha256"
	"encoding/base64"
	"github.com/Technology-99/qxLib/qxCrypto"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/hkdf"
	"time"
)

const (
	QxParserStatusNotReady = iota + 1
	QxParserStatusReady
)

type (
	QxParser interface {
		Init(localPrivateKey *ecdh.PrivateKey, remotePublicKey *ecdh.PublicKey, expireAt time.Time) error
		Status() int
		FormatStatus() string
		ExpireAt() time.Time
		Decrypt(data string) ([]byte, error)
		Encrypt(data string) (string, error)
	}
	defaultQxParser struct {
		sessionId        string
		localPrivateKey  *ecdh.PrivateKey
		remotePublicKey  *ecdh.PublicKey
		sharedSecretBase string
		deriveAESKeyBase string
		deriveAESIvBase  string
		status           int
		expireAt         time.Time
	}
)

func NewQxParser() QxParser {
	return &defaultQxParser{
		status: QxParserStatusNotReady,
	}
}

func (m *defaultQxParser) ExpireAt() time.Time {
	return m.expireAt
}

func (m *defaultQxParser) FormatStatus() string {
	if time.Now().Unix() > m.expireAt.Unix() {
		m.status = QxParserStatusNotReady
	}
	switch m.status {
	case QxParserStatusNotReady:
		return "解析器未就绪"
	case QxParserStatusReady:
		return "解析器就绪"
	default:
		return "未知状态"
	}
}

func (m *defaultQxParser) Status() int {
	if time.Now().Unix() > m.expireAt.Unix() {
		m.status = QxParserStatusNotReady
	}
	return m.status
}

func (m *defaultQxParser) Init(localPrivateKey *ecdh.PrivateKey, remotePublicKey *ecdh.PublicKey, expireAt time.Time) error {
	// 计算共享密钥
	sharedSecret, err := localPrivateKey.ECDH(remotePublicKey)
	if err != nil {
		logx.Errorf("keyexchange compute share key error: %v", err)
		return err
	}
	// 通过 HKDF 派生 AES-256 密钥
	hkdf := hkdf.New(sha256.New, sharedSecret, nil, []byte("key-session-key"))
	derivedKey := make([]byte, qxCrypto.AES256KeyLen) // 32 字节 AES-256 密钥
	hkdf.Read(derivedKey)
	m.localPrivateKey = localPrivateKey
	m.remotePublicKey = remotePublicKey
	m.sharedSecretBase = base64.StdEncoding.EncodeToString(sharedSecret)
	m.deriveAESKeyBase = base64.StdEncoding.EncodeToString(derivedKey)
	m.deriveAESIvBase = base64.StdEncoding.EncodeToString(sharedSecret[0:qxCrypto.AESGCMIvLen])
	m.status = QxParserStatusReady
	m.expireAt = expireAt
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
