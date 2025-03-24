package qxCcs

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
)

type (
	CcsService struct {
		CcsBaseService
	}
)

func NewCcsService(qxCtx *qxCtx.QxCtx) CcsService {
	return CcsService{
		CcsBaseService: NewCcsBaseService(qxCtx),
	}
}
