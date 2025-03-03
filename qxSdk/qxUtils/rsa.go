package qxUtils

import (
	"encoding/json"
	utils "github.com/Technology-99/third_party/cryptography"
)

func RsaEncrypt(input interface{}, PublicKey string) (*string, error) {
	btResult, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	pubKey, err := utils.ParsePublicKey(PublicKey)
	if err != nil {
		return nil, err
	}
	ncryptedText, err := utils.EncryptRSA(btResult, pubKey)
	if err != nil {
		return nil, err
	}

	return &ncryptedText, nil
}

func RsaDecrypt(encryptedBase64, privateKey string) ([]byte, error) {
	priKey, err := utils.ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	decodeBt, err := utils.DecryptRSA(encryptedBase64, priKey)
	if err != nil {
		return nil, err
	}

	return decodeBt, nil
}
