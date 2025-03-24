package qxMas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	MasService struct {
		MasBaseService
	}
)

func NewMasService(qxCtx *qxCtx.QxCtx) MasService {
	return MasService{
		MasBaseService: NewMsgBaseService(qxCtx),
	}
}
