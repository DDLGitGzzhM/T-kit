package error

import (
	"testing"

	"github.com/stretchr/testify/assert"

	tkcontext "T-kit/pkg/context"
)

func TestTError_Error(t *testing.T) {
	err := TError{
		TErrorCode: CommonResp,
		message:    "something went wrong",
	}

	assert.Equal(t, "something went wrong", err.Error())
}

func TestTError_ErrorForDisplay(t *testing.T) {
	err := TError{
		TErrorCode: CommonResp,
		message:    "something went wrong",
	}

	// 测试中文显示
	zhCtx := tkcontext.TContext{Lang: "zh"}
	assert.Equal(t, "普通错误:something went wrong", err.ErrorForDisplay(zhCtx))

	// 测试英文显示
	enCtx := tkcontext.TContext{Lang: "en"}
	assert.Equal(t, "common error:something went wrong", err.ErrorForDisplay(enCtx))

	// 测试空语言默认为英文
	emptyCtx := tkcontext.TContext{}
	assert.Equal(t, "common error:something went wrong", err.ErrorForDisplay(emptyCtx))
}
