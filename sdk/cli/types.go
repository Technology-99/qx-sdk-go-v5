package cli

const (
	KeyRequestId = "requestId"
)

type CommonApiResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestID"`
	Path      string `json:"path"`
}
