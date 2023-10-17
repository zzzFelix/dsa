package sort

func quicksort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	pivotIndex := len(input) / 2
	pivotElement := input[pivotIndex]

	lessThan := []int{}
	equalTo := []int{}
	greaterThan := []int{}

	for _, val := range input {
		if val == pivotElement {
			equalTo = append(equalTo, val)
		} else if val > pivotElement {
			greaterThan = append(greaterThan, val)
		} else {
			lessThan = append(lessThan, val)
		}
	}

	return append(append(quicksort(lessThan), equalTo...), quicksort(greaterThan)...)
}
