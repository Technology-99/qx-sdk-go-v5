package qxCtas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
)

type (
	CtasService struct {
		CtasBaseService
	}
)

func NewCtasService(cli *qxCli.QxClient) CtasService {
	return CtasService{
		CtasBaseService: NewCtasBaseService(cli),
	}
}
