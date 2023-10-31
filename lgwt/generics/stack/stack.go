package stack

// Stack represents a LIFO Abstract Data Type of any value
type Stack[T any] struct {
	b []T
}

// Push stores values at the top of the Stack
func (s *Stack[T]) Push(value T) {
	s.b = append(s.b, value)
}

// Pop retrieves and delete the value at the top of the Stack
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	top := len(s.b) - 1
	v := s.b[top]
	s.b = s.b[:top]
	return v, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.b) == 0
}
