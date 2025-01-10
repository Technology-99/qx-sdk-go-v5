package main

import (
	qxSdk "github.com/Technology-99/qx-sdk-go-v5/sdk"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {

	Endpoint := os.Getenv("ENDPOINT")
	logx.Infof("打印sdk接入点: %s", Endpoint)

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	s := qxSdk.NewSdk(AccessKeyId, AccessKeySecret, Endpoint)
	logx.Infof("打印sdk版本号: %s", s.GetVersion())

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
	//select {}
}
