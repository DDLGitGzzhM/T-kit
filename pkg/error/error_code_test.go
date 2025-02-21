package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTErrorCode_GetMessage(t *testing.T) {
	errorCode := NewTErrorCode("TEST", "测试", "test")

	// 测试中文消息
	assert.Equal(t, "测试", errorCode.GetMessage("zh"))

	// 测试英文消息
	assert.Equal(t, "test", errorCode.GetMessage("en"))

	// 测试未知语言默认返回英文
	assert.Equal(t, "test", errorCode.GetMessage("unknown"))
}

func TestNewTErrorCode(t *testing.T) {
	errorCode := NewTErrorCode("TEST", "测试", "test")

	assert.Equal(t, "TEST", errorCode.code)
	assert.Equal(t, "测试", errorCode.zhMsg)
	assert.Equal(t, "test", errorCode.enMsg)
}

func TestCommonResp(t *testing.T) {
	// 测试预定义的错误码
	assert.Equal(t, "COMMON_ERROR", CommonResp.code)
	assert.Equal(t, "普通错误", CommonResp.zhMsg)
	assert.Equal(t, "common error", CommonResp.enMsg)
}
