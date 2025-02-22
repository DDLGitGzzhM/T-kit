package slice

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShrink(t *testing.T) {
	smallSlice := make([]int, 10, 50)
	result := Shrink(smallSlice)
	assert.Equal(t, cap(smallSlice), cap(result), "small slice should not shrink")

	mediumSlice := make([]int, 100, 1000)
	result = Shrink(mediumSlice)
	assert.Equal(t, 500, cap(result), "medium slice with good ratio should not shrink")

	largeSlice := make([]int, 100, 4096)
	result = Shrink(largeSlice)
	expectedCap := int(math.Trunc(4096 * 0.6))
	assert.Equal(t, expectedCap, cap(result), "large slice with low usage should shrink to 60%")
}

func TestCalcShrinkCap(t *testing.T) {
	tests := map[string]struct {
		length   int
		capacity int
		want     int
	}{
		"small capacity": {
			length:   10,
			capacity: 64,
			want:     0,
		},
		"medium capacity with good ratio": {
			length:   1000,
			capacity: 2000,
			want:     0,
		},
		"large capacity with low usage": {
			length:   1000,
			capacity: 4096,
			want:     2457, // 4096 * 0.6
		},
		"very large capacity": {
			length:   2000,
			capacity: 10000,
			want:     6000, // 10000 * 0.6
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := calcShrinkCap(tc.length, tc.capacity)
			assert.Equal(t, tc.want, got)
		})
	}
}
