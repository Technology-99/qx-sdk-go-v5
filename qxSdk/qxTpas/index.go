package qxTpas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTpas/wechat"
)

type (
	TpasService struct {
		TpasBaseService
		Wechat Wechat
	}
	Wechat struct {
		CommonService      wechat.CommonService
		OffiaccountService wechat.OffiaccountService
	}
)

func NewTpasService(qxCtx *qxCtx.QxCtx) TpasService {
	return TpasService{
		TpasBaseService: NewTpasBaseService(qxCtx),
		Wechat: Wechat{
			CommonService:      wechat.NewCommonService(qxCtx),
			OffiaccountService: wechat.NewOffiaccountService(qxCtx),
		},
	}
}
