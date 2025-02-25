package qxStorage

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
)

type (
	StorageService struct {
		FileService FileService
		StorageBaseService
	}
)

func NewStorageService(cli *qxCli.QxClient) StorageService {
	return StorageService{
		FileService:        NewFileService(cli),
		StorageBaseService: NewStorageBaseService(cli),
	}
}
