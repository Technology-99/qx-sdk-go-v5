package qxMas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
)

type (
	MasService struct {
		MasBaseService
	}
)

func NewMasService(cli *qxCli.QxClient) MasService {
	return MasService{
		MasBaseService: NewMsgBaseService(cli),
	}
}
