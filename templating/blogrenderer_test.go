package templating_test

import (
	"bytes"
	"hello/templating"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aSong = templating.Song{
			Title:  "hello world",
			Lyrics: "This is a song",
		}
	)

	t.Run("it converts a single song into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := templating.Render(&buf, aSong)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got '%s' but wanted '%s'", got, want)
		}
	})
}
