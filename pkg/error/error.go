package error

import (
	"fmt"

	tkcontext "T-kit/pkg/context"
)

type TError struct {
	TErrorCode TErrorCode
	message    string
}

func (e TError) Error() string {
	return e.message
}

func (e TError) ErrorForDisplay(ctx tkcontext.TContext) string {
	if ctx.Lang == "" {
		ctx.Lang = "en"
	}
	return fmt.Sprintf("%s:%s", e.TErrorCode.GetMessage(ctx.Lang), e.message)
}
