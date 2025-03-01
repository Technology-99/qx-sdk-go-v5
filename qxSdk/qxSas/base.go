package qxSas

import "github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"

type (
	SasBaseService interface {
	}
	defaultSasBaseService struct {
		cli *qxCli.QxClient
	}
)

func NewSasBaseService(cli *qxCli.QxClient) SasBaseService {
	return &defaultSasBaseService{
		cli: cli,
	}
}
