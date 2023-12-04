package compare

// negative, if a < b
// zero    , if a == b
// positive, if a > b
type Comparator[T any] interface {
	Compare(a, b T) int
}

type IntComparator struct{}

func (c IntComparator) Compare(a, b int) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}
