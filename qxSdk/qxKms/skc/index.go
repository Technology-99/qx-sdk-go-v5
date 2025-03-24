package skc

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	KmsSkcService interface {
	}

	defaultKmsSkcService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewKmsSkcService(qxCtx *qxCtx.QxCtx) KmsSkcService {
	// note: 初始化Kms系统
	return &defaultKmsSkcService{
		qxCtx: qxCtx,
	}
}
