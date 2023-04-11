package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("wanted %.2f, but got %.2f", want, got)
	}
}
