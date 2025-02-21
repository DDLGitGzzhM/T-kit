package slice

import (
	"fmt"

	TE "T-kit/pkg/error"
)

func DeleteByIndex[S ~[]E, E any](s S, index int) (S, *TE.TError) {
	if index < 0 || index >= len(s) {
		return s, TE.NewTError(TE.OutOfIndex, fmt.Sprintf("index out of range : %d", index))
	}
	length := len(s)
	for i := index; i+1 < length; i++ {
		s[i] = s[i+1]
	}
	s = s[:length-1]
	return s, nil
}
