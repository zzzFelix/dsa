package collection

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	t.Run("should insert element into heap", func(t *testing.T) {
		// given
		testS := "Hello world"
		heap, _ := NewMaxHeap[string]([]string{testS}, []int{1})

		// when
		heap.Insert(testS, 99)

		// then
		actual := heap.Peek()
		if testS != actual {
			t.Fatal("Wrong element returned")
		}
		if heap.Size() != 2 {
			t.Fatalf("Wrong heap size. Expected: %d, actual: %d", 2, heap.Size())
		}
	})

	t.Run("should create heap with 5 elements and return highest priority", func(t *testing.T) {
		// given

		// when
		heap, _ := NewMaxHeap[string](
			[]string{"low prio", "low prio", "low prio", "low prio", "HIGH PRIO", "low prio", "low prio"},
			[]int{12, 44, 103, 12, 99999, 12, 13},
		)

		// then
		actual := heap.Peek()
		if actual != "HIGH PRIO" {
			t.Fatalf("Wrong element returned. Actual: %s", actual)
		}
		if heap.Size() != 7 {
			t.Fatalf("Wrong heap size. Expected: %d, actual: %d", 7, heap.Size())
		}
	})

	t.Run("should add 5 elements and return highest priority", func(t *testing.T) {
		// given
		heap, _ := NewMaxHeap[string]([]string{}, []int{})

		// when
		heap.Insert("low prio", 12)
		heap.Insert("low prio", 44)
		heap.Insert("low prio", 103)
		heap.Insert("low prio", 12)
		heap.Insert("HIGH PRIO", 999)
		heap.Insert("low prio", 12)
		heap.Insert("low prio", 13)

		// then
		actual := heap.Peek()
		if actual != "HIGH PRIO" {
			t.Fatalf("Wrong element returned. Actual: %s", actual)
		}
		if heap.Size() != 7 {
			t.Fatalf("Wrong heap size. Expected: %d, actual: %d", 7, heap.Size())
		}
	})

	t.Run("should delete element", func(t *testing.T) {
		// given
		testS := "Hello world"
		heap, _ := NewMaxHeap[string]([]string{testS}, []int{1})

		// when
		heap.Delete()

		// then
		if heap.Size() != 0 {
			t.Fatalf("Wrong heap size. Expected: %d, actual: %d", 0, heap.Size())
		}
	})

	t.Run("should delete all but one element, last element should be lowest prio", func(t *testing.T) {
		// given
		heap, _ := NewMaxHeap[string](
			[]string{"default prio", "default prio", "default prio", "LOWEST PRIO", "default prio", "default prio", "default prio"},
			[]int{12, 44, 103, 0, 99999, 12, 13},
		)

		// when
		heapSize := heap.Size() - 1
		for i := 0; i < heapSize; i++ {
			heap.Delete()
		}

		// then
		actual := heap.Peek()
		if actual != "LOWEST PRIO" {
			t.Fatalf("Wrong element returned. Actual: %s", actual)
		}
		if heap.Size() != 1 {
			t.Fatalf("Wrong heap size. Expected: %d, actual: %d", 1, heap.Size())
		}
	})
}
