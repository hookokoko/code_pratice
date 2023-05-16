package common

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Pop() T {
	var item T
	if len(s.items) == 0 {
		return item
	}
	item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Top() T {
	var item T
	if len(s.items) == 0 {
		return item
	}
	item = s.items[len(s.items)-1]
	return item
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
