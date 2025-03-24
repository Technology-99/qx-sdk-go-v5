package main

import (
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxConfig"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")
	logx.Infof("打印sdk接入点: %s", Endpoint)

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	logx.Infof("打印sdk接入ID: %s", AccessKeyId)
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	c := qxConfig.DefaultConfig(AccessKeyId, AccessKeySecret, Endpoint)

	s := qxSdk.NewQxSdk(c)
	logx.Infof("打印sdk版本号: %s, tmp: %v", s.GetVersion(), s.Cli.Config)

	//note: rsa消息通讯自动解密
	//msgResult, err := s.CcsService.TestMsg(context.Background(), &qxTypesCcs.CcsTestMsgReq{
	//	Key: "latest",
	//	Msg: "我来试试加密和解密",
	//})
	//if err != nil {
	//	logx.Errorf("发送消息失败: %v", err)
	//	return
	//}
	//
	//logx.Infof("打印消息结果: %s", msgResult)

	// note: 生成验证码测试
	//genCodeResult, err := s.MsgService.CaptchaGenerate(context.Background(), &msg.ApiCaptchaGenerateReq{Key: "default"})
	//if err != nil {
	//	logx.Errorf("生成验证码失败: %v", err)
	//	return
	//}
	//logx.Infof("打印生成结果: %s", genCodeResult.Data.Img)

	// note: 快速通过OssV4前端直传上传文件
	//result, err := s.FileService.CreateWithOssV4FrontUpload(context.Background(), &types.AllowCreateModelTmsFileWithFrontedUpload{
	//	Key:      "default",
	//	Bucket:   "oss-tid1-test",
	//	FileName: "aa.docx",
	//	FileSize: 10842986,
	//	MimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	//})
	//if err != nil {
	//	logx.Errorf("生成验证码失败: %v", err)
	//	return
	//}
	//logx.Infof("打印生成结果: %s", result.Data.SignUrl)

	//time.Sleep(time.Second * 30)
	//
	//// note: 摧毁sdk，释放资源
	//s.Destroy()
	select {}
}
