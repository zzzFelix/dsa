package main

import (
	"dsa/collection"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("should instantiate empy set", func(t *testing.T) {
		// given

		// when
		set := make(collection.Set[struct{}])

		// then
		size := set.Size()
		if size != 0 {
			t.Fatalf("New set is not empty but contains %v items", size)
		}
	})
}
