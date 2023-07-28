package cached

import (
	"time"
)

func NewCached[T any](data T, expiredAt time.Time) *Cached[T] {
	return &Cached[T]{
		data:      data,
		expiredAt: expiredAt,
	}
}

type Cached[T any] struct {
	data      T
	expiredAt time.Time
}

func (o Cached[T]) Expired(now time.Time) bool {
	return o.expiredAt.After(now)
}

func (o Cached[T]) Data() T {
	return o.data
}
