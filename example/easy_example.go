package main

import (
	"context"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxConfig"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxLang"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes"
	"github.com/Technology-99/qx-sdk-go-v5/qxSdk/qxTypes/qxTypesKms"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"time"
)

func main() {
	Endpoint := os.Getenv("ENDPOINT")
	logx.Infof("打印sdk接入点: %s", Endpoint)

	// 本示例从环境变量中获取AccessKey ID和AccessKey Secret。
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	logx.Infof("打印sdk接入ID: %s", AccessKeyId)
	AccessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

	c := qxConfig.DefaultConfig(AccessKeyId, AccessKeySecret, Endpoint)
	//c.Debug = true

	s := qxSdk.NewQxSdk(c)
	logx.Infof("打印sdk版本号: %s, sdk状态: %v", s.GetVersion(), s.FormatQxSdkStatus())

	time.Sleep(time.Second * 5)
	logx.Infof("打印sdk版本号: %s, sdk状态: %v", s.GetVersion(), s.FormatQxSdkStatus())

	//codes, err := s.Codes(context.Background(), &qxTypes.CodesReq{Lang: qxLang.LangEnUS, Svc: "QxEngine"})
	//if err != nil {
	//	logx.Errorf("测试获取错误码表失败: %v", err)
	//	return
	//}
	//for _, code := range codes.Data.List {
	//	logx.Infof("code: %v, num: %v, msg: %v", code.Code, code.Num, code.Msg)
	//}

	zones, err := s.Zones(context.Background(), &qxTypes.ZonesReq{Lang: qxLang.LangEnUS})
	if err != nil {
		logx.Errorf("测试获取区号表失败: %v", err)
		return
	}
	for _, code := range zones.Data.List {
		logx.Infof("code: %v, label: %v, area: %v", code.Code, code.Label, code.Area)
	}

	//note: ecdhe 密钥交换通信加密
	msgResult, err := s.KmsService.TestMsg(context.Background(), &qxTypesKms.KmsTestMsgReq{
		Msg: "我来试试加密和解密",
	})
	if err != nil {
		logx.Errorf("测试传输加密失败: %v", err)
		return
	}
	logx.Infof("传输解密结果: %v", msgResult.Data)
	//
	////note: 数据加密
	//encryptResult, err := s.KmsService.Skc.KmsSkcEncrypt(context.Background(), &qxTypesKms.KmsSkcEncryptReq{
	//	Name:     "id-qx-cas-key-001",
	//	BaseData: base64.StdEncoding.EncodeToString([]byte("华仔最帅，帅到爆炸")),
	//})
	//if err != nil {
	//	logx.Errorf("测试数据加密失败: %v", err)
	//	return
	//}
	//logx.Infof("测试数据加密结果: %v", encryptResult.Data)
	//
	//// note: 数据解密
	//decryptResult, err := s.KmsService.Skc.KmsSkcDecrypt(context.Background(), &qxTypesKms.KmsSkcDecryptReq{
	//	Name:     "id-qx-cas-key-001",
	//	BaseData: encryptResult.Data.BaseData,
	//})
	//if err != nil {
	//	logx.Errorf("测试数据解密失败: %v", err)
	//	return
	//}
	//realData, _ := base64.StdEncoding.DecodeString(decryptResult.Data.BaseData)
	//logx.Infof("测试数据解密结果: %v", string(realData))

	//
	//for true {
	//	time.Sleep(time.Second * 5)
	//	s.TestMsg()
	//}

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
