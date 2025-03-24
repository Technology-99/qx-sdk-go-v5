package qxKms

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	KmsService struct {
		KmsBaseService
	}
)

func NewKmsService(qxCtx *qxCtx.QxCtx) KmsService {
	return KmsService{
		KmsBaseService: NewKmsBaseService(qxCtx),
	}
}
