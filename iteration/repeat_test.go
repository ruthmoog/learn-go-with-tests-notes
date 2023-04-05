package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("c", 5)
	want := "ccccc"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func ExampleRepeat() {
	result := Repeat("F", 3)
	fmt.Println(result)
	// Output: FFF
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("c", 5)
	}
}
