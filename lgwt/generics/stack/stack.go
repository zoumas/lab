package stack

import "errors"

var PopErr = errors.New("could not pop element because stack is empty")

type Stack[T any] struct {
	buf []T
}

func (s *Stack[T]) Push(element T) {
	s.buf = append(s.buf, element)
}

func (s *Stack[T]) Pop() (element T, err error) {
	top := len(s.buf) - 1
	if top < 0 {
		err = PopErr
		return
	}
	element = s.buf[top]
	s.buf = s.buf[:top]
	return
}

func (s *Stack[T]) Peek() (element T) {
	top := len(s.buf) - 1
	if top < 0 {
		return
	}
	return s.buf[top]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.buf) == 0
}
