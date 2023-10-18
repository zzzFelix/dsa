package sort

// negative, if a < b
// zero    , if a == b
// positive, if a > b
type Comparator[T interface{}] func(a, b T) int

func IntComparator(a, b int) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func quicksort[T interface{}, U func(a, b T) int](input []T, comp U) []T {
	if len(input) < 2 {
		return input
	}
	pivotIndex := len(input) / 2
	pivotElement := input[pivotIndex]

	lessThan := []T{}
	equalTo := []T{}
	greaterThan := []T{}

	for _, val := range input {
		if comp(val, pivotElement) == 0 {
			equalTo = append(equalTo, val)
		} else if comp(val, pivotElement) > 0 {
			greaterThan = append(greaterThan, val)
		} else {
			lessThan = append(lessThan, val)
		}
	}

	return append(append(quicksort(lessThan, comp), equalTo...), quicksort(greaterThan, comp)...)
}
