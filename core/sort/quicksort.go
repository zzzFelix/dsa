package sort

import "dsa/core/compare"

func quicksort[T any](input []T, comp compare.Comparator[T]) []T {
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
