package sort

// negative, if a < b
// zero    , if a == b
// positive, if a > b
type Comparator interface {
	Compare(a, b interface{}) int
}

type IntComparator struct{}

func (t IntComparator) Compare(a, b interface{}) int {
	intA, okA := a.(int)
	intB, okB := b.(int)
	if !okA || !okB {
		panic("Both parameters must be of type int")
	}
	switch {
	case intA > intB:
		return 1
	case intA < intB:
		return -1
	default:
		return 0
	}
}

func quicksort[T any](input []T, comp Comparator) []T {
	var comp2 Comparator = comp
	if len(input) < 2 {
		return input
	}
	pivotIndex := len(input) / 2
	pivotElement := input[pivotIndex]

	lessThan := []T{}
	equalTo := []T{}
	greaterThan := []T{}

	for _, val := range input {
		if comp2.Compare(val, pivotElement) == 0 {
			equalTo = append(equalTo, val)
		} else if comp2.Compare(val, pivotElement) > 0 {
			greaterThan = append(greaterThan, val)
		} else {
			lessThan = append(lessThan, val)
		}
	}

	return append(append(quicksort(lessThan, comp), equalTo...), quicksort(greaterThan, comp)...)
}
