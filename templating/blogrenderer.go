package templating

import (
	"html/template"
	"io"
)

type Song struct {
	Title  string
	Lyrics string
}

const songTemplate = `<h1>{{.Title}}</h1><p>{{.Lyrics}}</p>`

func Render(writer io.Writer, song Song) error {
	templ, err := template.New("Song").Parse(songTemplate)

	if err != nil {
		return err
	}

	if err := templ.Execute(writer, song); err != nil {
		return err
	}

	return nil
}
