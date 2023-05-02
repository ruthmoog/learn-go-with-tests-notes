package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 30},
			[]string{"Chris"},
		},
		{"nested fields",
			Person{
				"Chris",
				Profile{30, "Ankara"},
			},
			[]string{"Chris", "Ankara"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{30, "Ankara"},
			},
			[]string{"Chris", "Ankara"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Rome"},
			},
			[]string{"London", "Rome"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Pig": "Oink",
			"Cow": "Moo",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Oink")
		assertContains(t, got, "Moo")
	})
}

func assertContains(tb testing.TB, haystack []string, needle string) {
	tb.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		tb.Errorf("expected %+v to contain %q but it didnt!", haystack, needle)
	}
}
