package error

// TErrorCode 错误码结构体
type TErrorCode struct {
	code  TErrorCodeStr
	zhMsg string
	enMsg string
}

func NewTErrorCode(code TErrorCodeStr, zhMsg, enMsg string) TErrorCode {
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
	CommonResp = NewTErrorCode(CommonError, "普通错误", "common error")
	OutOfIndex = NewTErrorCode(OutOfIndexError, "超出范围", "out of index")
)

type TErrorCodeStr string

func (t TErrorCodeStr) Str() string { return string(t) }

const (
	CommonError     TErrorCodeStr = "COMMON_ERROR"
	OutOfIndexError TErrorCodeStr = "OUT_OF_INDEX_ERROR"
)
