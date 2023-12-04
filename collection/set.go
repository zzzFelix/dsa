package collection

type Set[T comparable] map[T]struct{}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Insert(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Delete(item T) {
	delete(s, item)
}

func (s Set[T]) Size() int {
	return len(s)
}
