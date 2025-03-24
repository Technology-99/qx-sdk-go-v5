package qxSas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	SasBaseService interface {
	}
	defaultSasBaseService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewSasBaseService(qxCtx *qxCtx.QxCtx) SasBaseService {
	return &defaultSasBaseService{
		qxCtx: qxCtx,
	}
}
