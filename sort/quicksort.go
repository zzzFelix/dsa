package sort

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//
//	negative , if a < b
//	zero     , if a == b
//	positive , if a > b
type Comparator func(a, b interface{}) int

func quicksort[T interface{}](input []T, comp Comparator) []T {
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
