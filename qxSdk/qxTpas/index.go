package qxTpas

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
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

func NewTpasService(cli *qxCli.QxClient) TpasService {
	return TpasService{
		TpasBaseService: NewTpasBaseService(cli),
		Wechat: Wechat{
			CommonService:      wechat.NewCommonService(cli),
			OffiaccountService: wechat.NewOffiaccountService(cli),
		},
	}
}
