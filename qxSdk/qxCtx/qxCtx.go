package qxCtx

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxParser"
)

type QxCtx struct {
	Cli    *qxCli.QxClient
	Parser qxParser.QxParser
}

func NewQxCtx(cli *qxCli.QxClient) *QxCtx {
	return &QxCtx{
		Cli:    cli,
		Parser: qxParser.NewQxParser(),
	}
}
