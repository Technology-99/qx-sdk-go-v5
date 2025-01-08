package main

import (
	"github.com/Technology-99/qx-sdk-go-v5/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")
	logx.Infof("打印sdk接入点: %s", Endpoint)

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := sdk.NewSdk(AccessKeyId, AccessKeySecret, Endpoint)

	// note: 生成验证码测试
	//s.MsgService.CaptchaGenerate(context.Background(), nil)
	logx.Infof("打印sdk版本号: %s", s.GetVersion())

	time.Sleep(time.Second * 30)

	// note: 摧毁sdk，释放资源
	s.Destroy()
	select {}
}
