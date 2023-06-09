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

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Astana"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Astana"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, but got %v", want, got)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Katowice"}, Profile{33, "Addis Ababa"}
		}

		var got []string
		want := []string{"Katowice", "Addis Ababa"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, but got %v", want, got)
		}
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
