package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("assert on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("assert on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "goodbye")
	})

	//AssertEqual(t, 1, "1") // uncomment to see the error
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(99)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(1)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 1)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 99)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(StackOfStrings)

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("Hello")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("World")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "World")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "Hello")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})

	t.Run("interface stack DX is horrid", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()

		// handle "operator `+` not defined on
		// firstNum (variable of type interface{})"
		// by getting the ints out of the interface...
		firstNumOfType, _ := firstNum.(int)
		secondNumOfType, _ := secondNum.(int)

		AssertEqual(t, firstNumOfType+secondNumOfType, 3)
	})
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

//func AssertEqual[T comparable](t *testing.T, got, want T) {
//	t.Helper()
//	if got != want {
//		t.Errorf("got %v, want %v", got, want)
//	}
//}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
