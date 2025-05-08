package qxCtx

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
)

type QxCtx struct {
	Cli *qxCli.QxClient
}

func NewQxCtx(cli *qxCli.QxClient) *QxCtx {
	return &QxCtx{
		Cli: cli,
	}
}
