package qxCms

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
)

type (
	CmsService struct {
		CmsBaseService
	}
)

func NewCmsService(cli *qxCli.QxClient) CmsService {
	return CmsService{
		CmsBaseService: NewCmsBaseService(cli),
	}
}
