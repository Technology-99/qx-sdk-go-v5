package qxCtas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	CtasService struct {
		CtasBaseService
	}
)

func NewCtasService(qxCtx *qxCtx.QxCtx) CtasService {
	return CtasService{
		CtasBaseService: NewCtasBaseService(qxCtx),
	}
}
