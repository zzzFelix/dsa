package leetcode

import "testing"

func TestFirstBadVersion(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		// given
		n := 5
		badVersion := 4

		// when
		actual := firstBadVersion(n, badVersion)

		// then
		if badVersion != actual {
			t.Fatalf("did not find first bad version. expected: %d, actual: %d", badVersion, actual)
		}
	})

	t.Run("test case 1", func(t *testing.T) {
		// given
		n := 1
		badVersion := 1

		// when
		actual := firstBadVersion(n, badVersion)

		// then
		if badVersion != actual {
			t.Fatalf("did not find first bad version. expected: %d, actual: %d", badVersion, actual)
		}
	})
}
