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
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}
