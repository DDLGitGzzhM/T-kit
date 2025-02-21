package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	TE "T-kit/pkg/error"
)

func Test_DeleteByIndex(t *testing.T) {
	slice := []string{"a", "b", "c"}
	res, Terr := DeleteByIndex(slice, 1)
	require.Nil(t, Terr)
	assert.Equal(t, res, []string{"a", "c"})
	res, Terr = DeleteByIndex(res, 3)
	assert.NotNil(t, Terr)
	assert.Equal(t, Terr.TErrorCode, TE.OutOfIndex)
}
