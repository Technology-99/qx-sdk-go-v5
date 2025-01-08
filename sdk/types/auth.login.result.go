package types

import "time"

type AuthLoginVO struct {
	UserId              uint32     `json:"userId"`
	Account             string     `json:"account"`
	TrueName            string     `json:"trueName"`
	Mobile              string     `json:"mobile"`
	IDentity            int32      `json:"status"`
	LastLogin           *time.Time `json:"lastLogin"`
	ThisLogin           *time.Time `json:"thisLogin"`
	AccessKey           string     `json:"accessKey"`
	AccessToken         string     `json:"accessToken"`
	AccessTokenExpires  int64      `json:"accessTokenExpires"`
	RefreshToken        string     `json:"refreshToken"`
	RefreshTokenExpires int64      `json:"refreshTokenExpires"`
}

type RefreshTokenData struct {
	AccessKey          string `json:"accessKey"`
	AccessToken        string `json:"accessToken"`
	AccessTokenExpires int64  `json:"accessTokenExpires"`
}
