package msg

import "github.com/Technology-99/qx-sdk-go-v5/sdk/cli"

type ApiCaptchaGenerateReq struct {
	Key       string  `json:"key,optional"`
	DotCount  int32   `json:"dotCount,optional"`
	MaxSkew   float64 `json:"maxSkew,optional"`
	KeyLong   int32   `json:"keyLong,optional"`
	ImgWidth  int32   `json:"imgWidth,optional"`
	ImgHeight int32   `json:"imgHeight,optional"`
}

type ApiCaptchaGenerateResp struct {
	cli.CommonApiResp
	Data ApiCaptchaGenerateRespData `json:"data"`
}

type ApiCaptchaGenerateRespData struct {
	Key     string `json:"key,omitempty"`
	Id      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	Answer  string `json:"answer,omitempty"`
	Img     string `json:"img,omitempty"`
}
