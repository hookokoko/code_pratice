package common

type Queue[T any] struct {
	items []T
	size  int
}

type Options[T any] func(*Queue[T])

func NewQueue[T any](size int, opts ...Options[T]) *Queue[T] {
	q := Queue[T]{
		items: make([]T, 0, size),
	}
	for _, opt := range opts {
		opt(&q)
	}
	return &q
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Add(ele T) {
	q.items = append(q.items, ele)
}

func (q *Queue[T]) Pop() T {
	var ele T
	if !q.IsEmpty() {
		ele = q.items[0]
		q.items = q.items[1:]
	}
	return ele
}

func (q *Queue[T]) Top() T {
	var ele T
	if !q.IsEmpty() {
		ele = q.items[0]
	}
	return ele
}
