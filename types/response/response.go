package response

type Response struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	ErrStack interface{} `json:"errStack,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
