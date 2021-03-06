package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorNotFound    = ErrResponse{HttpSC: 404, Error: Err{Error: "很抱歉您访问的地址不存在", ErrorCode: -1}}
	ErrorNotMethod   = ErrResponse{HttpSC: 404, Error: Err{Error: "很抱歉您访问的方法不存在", ErrorCode: -1}}
	ErrorLostParams  = ErrResponse{HttpSC: 200, Error: Err{Error: "你把参数弄丢了", ErrorCode: -1}}
	MissingSignature = ErrResponse{HttpSC: 400, Error: Err{Error: "检查签名相关参数，可能已经过期", ErrorCode: 10000}}
	LoginFail        = ErrResponse{HttpSC: 200, Error: Err{Error: "登录失败", ErrorCode: 10001}}
	LoginAuthFail    = ErrResponse{HttpSC: 402, Error: Err{Error: "授权失败，请求重新验证身份", ErrorCode: 10002}}
)

func ValidateErr(msg string) ErrResponse {
	return ErrResponse{
		HttpSC: 200, Error: Err{
			Error:     msg,
			ErrorCode: -1,
		},
	}
}
