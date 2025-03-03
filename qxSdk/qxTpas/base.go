package qxTpas

import "github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"

type (
	TpasBaseService interface {
	}
	defaultTpasBaseService struct {
		cli *qxCli.QxClient
	}
)

func NewTpasBaseService(cli *qxCli.QxClient) TpasBaseService {
	return &defaultTpasBaseService{
		cli: cli,
	}
}
