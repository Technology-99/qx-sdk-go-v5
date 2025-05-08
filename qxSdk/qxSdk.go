package qxSdk

import (
	"context"
	"embed"
	"fmt"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxBase"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCcs"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCli"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxConfig"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxCtx"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxKms"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxMas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxSas"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTpas"
	"sync"
)

//go:embed "version"
var VersionF embed.FS

type QxSdk struct {
	Version string

	ctx        context.Context    // 控制退出
	cancel     context.CancelFunc // 取消函数
	wg         sync.WaitGroup     // 等待后台任务退出
	isShutdown bool               // 标记 SDK 是否已经关闭

	QxCtx *qxCtx.QxCtx

	// note: 琼霄自身的服务
	qxBase.QxBaseService

	// note: 云端配置存储服务
	CcsService qxCcs.CcsService
	// note: 消息服务
	MasService qxMas.MasService
	// note: 存储服务
	SasService qxSas.SasService
	// note: 定时任务队列服务
	CtasService qxCtas.CtasService
	// note: 第三方聚合服务
	TpasService qxTpas.TpasService
	// note: 密钥管理服务
	KmsService qxKms.KmsService
}

func NewQxSdk(c *qxConfig.Config) *QxSdk {
	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	qxClient := qxCli.NewQxClient(ctx, c)
	qxC := qxCtx.NewQxCtx(qxClient)

	sdk := &QxSdk{
		Version:       string(versionFile),
		QxCtx:         qxC,
		ctx:           ctx,
		cancel:        cancel,
		QxBaseService: qxBase.NewQxBaseService(qxC),
		CcsService:    qxCcs.NewCcsService(qxC),
		KmsService:    qxKms.NewKmsService(qxC),
		MasService:    qxMas.NewMasService(qxC),
		SasService:    qxSas.NewSasService(qxC),
		CtasService:   qxCtas.NewCtasService(qxC),
		TpasService:   qxTpas.NewTpasService(qxC),
	}
	return sdk
}

func NewDefaultQxSdk(AccessKeyId, AccessKeySecret, Endpoint string) *QxSdk {

	c := qxConfig.DefaultConfig(AccessKeyId, AccessKeySecret, Endpoint)
	versionFile, err := VersionF.ReadFile("version")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	qxClient := qxCli.NewQxClient(ctx, c)
	qxC := qxCtx.NewQxCtx(qxClient)

	sdk := &QxSdk{
		Version:     string(versionFile),
		QxCtx:       qxC,
		ctx:         ctx,
		cancel:      cancel,
		CcsService:  qxCcs.NewCcsService(qxC),
		KmsService:  qxKms.NewKmsService(qxC),
		MasService:  qxMas.NewMasService(qxC),
		SasService:  qxSas.NewSasService(qxC),
		CtasService: qxCtas.NewCtasService(qxC),
		TpasService: qxTpas.NewTpasService(qxC),
	}

	return sdk
}

func (s *QxSdk) Destroy() {
	if s.isShutdown {
		fmt.Println("SDK already shutdown.")
		return
	}

	fmt.Println("Shutting down SDK...")
	s.isShutdown = true

	// 通知后台任务退出
	s.cancel()

	// 等待所有 goroutine 退出
	s.wg.Wait()

	// 关闭 HTTP 客户端
	s.QxCtx.Cli.Client.CloseIdleConnections()

	fmt.Println("SDK resources released.")
}

func (s *QxSdk) GetVersion() string {
	return s.Version
}
