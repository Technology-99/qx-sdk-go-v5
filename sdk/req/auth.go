package req

type QxV5ApisApiSignReq struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

type QxV5ApisApiRefreshReq struct {
	AccessKey    string `json:"accessKey"`
	RefreshToken string `json:"refreshToken"`
}
