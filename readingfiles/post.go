package readingfiles

import (
	"io"
	"strings"
)

type Post struct {
	Title string
}

func newPost(blogFile io.Reader) Post {
	fileContents, _ := io.ReadAll(blogFile)
	title := strings.TrimPrefix(string(fileContents), "Title: ")
	return Post{
		Title: title,
	}
}
