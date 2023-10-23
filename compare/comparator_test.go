package compare

import "testing"

func TestIntComparator(t *testing.T) {
	var comparator Comparator[int] = IntComparator{}

	t.Run("should return 0 when given equal integers", func(t *testing.T) {
		// given
		a := 4
		b := 4

		// when
		actual := comparator.Compare(a, b)

		// then
		expected := 0
		if actual != expected {
			t.Fatalf("comparison didn't return %d. Actual: %d", expected, actual)
		}
	})

	t.Run("should return 1 when a > b", func(t *testing.T) {
		// given
		a := 999
		b := -999

		// when
		actual := comparator.Compare(a, b)

		// then
		expected := 1
		if actual != expected {
			t.Fatalf("comparison didn't return %d. Actual: %d", expected, actual)
		}
	})

	t.Run("should return -1 when a < b", func(t *testing.T) {
		// given
		a := 0
		b := 1

		// when
		actual := comparator.Compare(a, b)

		// then
		expected := -1
		if actual != expected {
			t.Fatalf("comparison didn't return %d. Actual: %d", expected, actual)
		}
	})
}
