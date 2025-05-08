package qxTypes

type AuthConfig struct {
	Auths map[string]Auth `json:"auths"`
}

type Auth struct {
	AccessToken     string `json:"accessToken"`
	RemotePublicKey string `json:"remotePublicKey"`
	LocalPrivateKey string `json:"localPrivateKey"`
	ExpiresIn       int64  `json:"expiresIn"`
}
