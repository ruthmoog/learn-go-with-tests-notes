package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("c")
	want := "ccccc"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}
