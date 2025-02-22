package list

import (
	"T-kit/internal/slice"
	TE "T-kit/pkg/error"
)

type ArrayList[T any] struct {
	vals []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{vals: make([]T, 0, cap)}
}

// Delete 方法会在必要时间引起缩容,其规则是:
// - 如果容量 [2048,INF] , 并且长度小于容量的一半, 容量会减少 3/5
// - 如果容量 (64,2048], 并且长度小于容量的4倍, 容量会减少 1/2
// - 如果容量 [0,64], 不会进行缩容
func (a *ArrayList[T]) Delete(index int) *TE.TError {
	res, tErr := slice.DeleteByIndex(a.vals, index)
	if tErr != nil {
		return tErr
	}
	a.vals = res
	a.shrink()
	return nil
}

func (a *ArrayList[T]) shrink() {
	a.vals = slice.Shrink(a.vals)
}
