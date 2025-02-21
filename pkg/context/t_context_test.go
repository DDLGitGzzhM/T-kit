package tkcontext

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTContext(t *testing.T) {
	ctx := context.Background()
	tCtx := NewTContext(ctx)

	assert.NotNil(t, tCtx)
	assert.IsType(t, &TContext{}, tCtx)
}
