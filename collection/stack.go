package collection

type stack[T any] struct {
	internal []T
}

type Stack[T any] stack[T]

func (s *Stack[T]) Pop() T {
	last := s.Peek()
	s.internal = s.internal[:len(s.internal)-1]
	return last
}

func (s *Stack[T]) Push(item T) {
	s.internal = append(s.internal, item)
}

func (s *Stack[T]) Peek() T {
	return s.internal[len(s.internal)-1]
}
