package qxCcs

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCcs/aes"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxParser"
)

type (
	CcsService struct {
		CcsBaseService
		AesService aes.AesService
	}
)

func NewCcsService(cli *qxCli.QxClient, parser qxParser.QxParser) CcsService {
	return CcsService{
		CcsBaseService: NewCcsBaseService(cli, parser),
		AesService:     aes.NewAesService(cli),
	}
}
