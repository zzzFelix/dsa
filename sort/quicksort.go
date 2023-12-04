package sort

import "github.com/zzzFelix/dsa/compare"

func Quicksort[T any](input []T, comp compare.Comparator[T]) []T {
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

	return append(append(Quicksort(lessThan, comp), equalTo...), Quicksort(greaterThan, comp)...)
}
