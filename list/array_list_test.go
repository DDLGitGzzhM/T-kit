package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Delete(t *testing.T) {
	testList := NewArrayList[int](0)
	testList.vals = []int{1, 2, 3}
	tErr := testList.Delete(1)
	require.Nil(t, tErr)
	t.Log(testList.vals)
}
