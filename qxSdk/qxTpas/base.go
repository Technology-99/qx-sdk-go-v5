package qxTpas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	TpasBaseService interface {
	}
	defaultTpasBaseService struct {
		qxCtx *qxCtx.QxCtx
	}
)

func NewTpasBaseService(qxCtx *qxCtx.QxCtx) TpasBaseService {
	return &defaultTpasBaseService{
		qxCtx: qxCtx,
	}
}
