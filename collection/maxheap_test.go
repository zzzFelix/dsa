package collection

import "testing"

func TestMaxHeap(t *testing.T) {
	t.Run("should insert element into heap", func(t *testing.T) {
		// given
		testS := "Hello world"
		heap, err := New[string]([]string{testS}, []int{1})

		if err != nil {
			t.Fatal("Heap creation resulted in error")
		}

		// when
		heap.Insert(testS, 99)

		// then
		actual := heap.Peek()
		if testS != actual {
			t.Fatal("Wrong element returned")
		}
	})
}
