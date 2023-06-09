package templating

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	songTemplates embed.FS
)

type Song struct {
	Title  string
	Lyrics string
}

func Render(writer io.Writer, song Song) error {
	templ, err := template.ParseFS(songTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(writer, "songbook.gohtml", song); err != nil {
		return err
	}

	return nil
}
