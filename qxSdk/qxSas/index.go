package qxSas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	SasService struct {
		FileService   FileService
		FolderService FolderService
		SasBaseService
	}
)

func NewSasService(qxCtx *qxCtx.QxCtx) SasService {
	return SasService{
		FileService:    NewFileService(qxCtx),
		FolderService:  NewFolderService(qxCtx),
		SasBaseService: NewSasBaseService(qxCtx),
	}
}
