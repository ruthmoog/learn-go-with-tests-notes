package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string
	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, expected %d but got %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("expected %q but got %q", expected, got[0])
	}
}
