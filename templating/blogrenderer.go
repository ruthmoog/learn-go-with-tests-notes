package templating

import (
	"fmt"
	"io"
)

type Song struct {
	Title  string
	Lyrics string
}

func Render(writer io.Writer, song Song) error {
	_, err := fmt.Fprintf(writer, "<h1>%s</h1>", song.Title)
	return err
}
