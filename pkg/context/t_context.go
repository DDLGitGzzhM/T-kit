package tkcontext

import "context"

type TContext struct {
	Context context.Context
	Lang    string
}

func NewTContext(ctx context.Context) *TContext {
	return &TContext{}
}
