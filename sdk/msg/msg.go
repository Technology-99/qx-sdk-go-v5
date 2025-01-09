package msg

import (
	"context"
	"encoding/json"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/cli"
	"github.com/Technology-99/qx-sdk-go-v5/sdk/types"
	"github.com/Technology-99/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type (
	MsgService interface {
		// note: 生成验证码
		CaptchaGenerate(ctx context.Context, params *types.ApiCaptchaGenerateReq) (result *types.ApiCaptchaGenerateResp, err error)
	}

	defaultMsgService struct {
		cli *cli.QxClient
	}
)

func NewMsgService(cli *cli.QxClient) MsgService {
	return &defaultMsgService{
		cli: cli,
	}
}

func (m *defaultMsgService) CaptchaGenerate(ctx context.Context, params *types.ApiCaptchaGenerateReq) (result *types.ApiCaptchaGenerateResp, err error) {
	result = &types.ApiCaptchaGenerateResp{}
	reqFn := m.cli.EasyNewRequest(ctx, "/captcha/generate", http.MethodPost, &params)
	res, err := reqFn()
	if err != nil {
		logx.Errorf("healthz request error: %v", err)
		return nil, nil
	}
	_ = json.Unmarshal(res, &result)
	if result.Code != response.SUCCESS {
		logx.Errorf("qiongxiao sdk errlog: captchaGenerate fail: %v", result)
		return result, nil
	}
	return result, nil
}
