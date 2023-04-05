package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("c")
	want := "ccccc"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("c2")
	}
}
