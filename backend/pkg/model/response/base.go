package response

type BaseResp struct {
	Code    int64  `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewErrorBaseResp(err string, code int64) *BaseResp {
	return &BaseResp{
		Error: err,
		Code:  code,
	}
}

func NewSucceedBaseResp(msg string) *BaseResp {
	return &BaseResp{
		Message: msg,
	}
}
