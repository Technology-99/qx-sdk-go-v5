package qxKms

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxKms/akc"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxKms/skc"
)

type (
	KmsService struct {
		KmsBaseService
		Akc akc.KmsAkcService
		Skc skc.KmsSkcService
	}
)

func NewKmsService(qxCtx *qxCtx.QxCtx) KmsService {
	return KmsService{
		KmsBaseService: NewKmsBaseService(qxCtx),
		Akc:            akc.NewKmsAkcService(qxCtx),
		Skc:            skc.NewKmsSkcService(qxCtx),
	}
}
