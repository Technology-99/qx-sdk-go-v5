package qxKms

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
)

type (
	KmsBaseService interface {
	}

	defaultKmsBaseService struct {
		cli *qxCli.QxClient
	}
)

func NewKmsBaseService(cli *qxCli.QxClient) KmsBaseService {
	// note: 初始化Kms系统
	return &defaultKmsBaseService{
		cli: cli,
	}
}
