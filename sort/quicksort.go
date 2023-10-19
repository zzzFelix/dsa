package sort

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

func quicksort[T any](input []T, comp Comparator[T]) []T {
	if len(input) < 2 {
		return input
	}
	pivotIndex := len(input) / 2
	pivotElement := input[pivotIndex]

	lessThan := []T{}
	equalTo := []T{}
	greaterThan := []T{}

	for _, val := range input {
		if comp.Compare(val, pivotElement) == 0 {
			equalTo = append(equalTo, val)
		} else if comp.Compare(val, pivotElement) > 0 {
			greaterThan = append(greaterThan, val)
		} else {
			lessThan = append(lessThan, val)
		}
	}

	return append(append(quicksort(lessThan, comp), equalTo...), quicksort(greaterThan, comp)...)
}
