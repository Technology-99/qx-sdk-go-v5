package qxSas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
)

type (
	SasService struct {
		FileService FileService
		SasBaseService
	}
)

func NewSasService(cli *qxCli.QxClient) SasService {
	return SasService{
		FileService:    NewFileService(cli),
		SasBaseService: NewSasBaseService(cli),
	}
}
