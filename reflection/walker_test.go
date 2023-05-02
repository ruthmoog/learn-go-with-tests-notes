package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "Ankara"},
			[]string{"Chris", "Ankara"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("expected %v but got %v", test.ExpectedCalls, got)
			}
		})
	}
}
