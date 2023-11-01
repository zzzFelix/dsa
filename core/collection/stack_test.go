package collection

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("should peek lasted pushed item", func(t *testing.T) {
		// given
		stack := Stack[string]{}
		expected := "world"

		// when
		stack.Push("hello")
		stack.Push(expected)

		// then
		actual := stack.Peek()
		if expected != actual {
			t.Fatalf("did not peek right item. expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("should peek first item when second was deleted", func(t *testing.T) {
		// given
		stack := Stack[string]{}
		expected := "hello"

		// when
		stack.Push(expected)
		stack.Push("world")
		stack.Pop()

		// then
		actual := stack.Peek()
		if expected != actual {
			t.Fatalf("did not peek right item. expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("should return item when popped", func(t *testing.T) {
		// given
		stack := Stack[string]{}
		expected := "hello"

		// when
		stack.Push(expected)
		actual := stack.Pop()

		// then
		if expected != actual {
			t.Fatalf("did not peek right item. expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("should push and pop the same item 10 times", func(t *testing.T) {
		// given
		stack := Stack[struct{}]{}

		// when
		for i := 0; i < 10; i++ {
			stack.Push(struct{}{})
		}
		for i := 0; i < 10; i++ {
			stack.Pop()
		}

		// then
		if stack.Size() != 0 {
			t.Fatal("stack should be empty")
		}
	})
}
