package exo

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, ErrEmptyStack
    }
    n := len(s.items) - 1
    item := s.items[n]
    s.items = s.items[:n]
    return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, ErrEmptyStack
    }
    return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) Size() int {
    return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}