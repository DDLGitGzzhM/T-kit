package error

// TErrorCode 错误码结构体
type TErrorCode struct {
	code  string
	zhMsg string
	enMsg string
}

func NewTErrorCode(code, zhMsg, enMsg string) TErrorCode {
	return TErrorCode{
		code:  code,
		zhMsg: zhMsg,
		enMsg: enMsg,
	}
}

func (e TErrorCode) GetMessage(lang string) string {
	if lang == "zh" {
		return e.zhMsg
	}
	return e.enMsg
}

var (
	CommonResp = NewTErrorCode("COMMON_ERROR", "普通错误", "common error")
)
