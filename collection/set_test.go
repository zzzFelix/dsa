package collection

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("should add item to set of strings", func(t *testing.T) {
		// given
		set := make(Set[string])
		testString := "Hello World"

		// when
		set.Insert(testString)

		// then
		if !set.Contains(testString) {
			t.Fatal("Set does not contain test string")
		} else if set.Size() != 1 {
			t.Fatal("Set is not of size 1. Actual size:", set.Size())
		}
	})

	t.Run("should delete item", func(t *testing.T) {
		// given
		testString := "Hello World"
		set := make(Set[string])
		set.Insert(testString)

		// when
		set.Delete(testString)

		// then
		if set.Contains(testString) {
			t.Fatal("Set still contains string that should be deleted")
		}
		if set.Size() != 0 {
			t.Fatal("Set is not of size 0. Actual size:", set.Size())
		}
	})

	t.Run("should add 3 items to set of strings", func(t *testing.T) {
		// given
		set := make(Set[string])
		testString := "Hello World"

		// when
		for i := 0; i < 3; i++ {
			set.Insert(fmt.Sprintf("%s %v", testString, i))
		}

		// then
		for i := 0; i < 3; i++ {
			if !set.Contains(fmt.Sprintf("%s %v", testString, i)) {
				t.Fatal("Set does not contain test string", i)
			}
		}
		if set.Size() != 3 {
			t.Fatal("Set is not of size 3. Actual size:", set.Size())
		}
	})

	t.Run("should not add duplicate to set", func(t *testing.T) {
		// given
		set := make(Set[string])
		testString := "Hello World"

		// when
		set.Insert(testString)
		set.Insert(testString)

		// then
		if !set.Contains(testString) {
			t.Fatal("Set does not contain test string")
		} else if set.Size() != 1 {
			t.Fatal("Set is not of size 1. Actual size:", set.Size())
		}
	})
}
