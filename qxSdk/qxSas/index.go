package qxSas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	SasService struct {
		FileService FileService
		SasBaseService
	}
)

func NewSasService(qxCtx *qxCtx.QxCtx) SasService {
	return SasService{
		FileService:    NewFileService(qxCtx),
		SasBaseService: NewSasBaseService(qxCtx),
	}
}
