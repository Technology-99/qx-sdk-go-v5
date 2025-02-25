package qxStorage

import "github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"

type (
	StorageBaseService interface {
	}
	defaultStorageBaseService struct {
		cli *qxCli.QxClient
	}
)

func NewStorageBaseService(cli *qxCli.QxClient) StorageBaseService {
	return &defaultStorageBaseService{
		cli: cli,
	}
}
