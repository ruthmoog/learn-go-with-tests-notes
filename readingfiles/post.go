package readingfiles

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)
	scanner.Scan()
	title := strings.TrimPrefix(scanner.Text(), "Title: ")

	scanner.Scan()
	description := strings.TrimPrefix(scanner.Text(), "Description: ")

	return Post{
		Title:       title,
		Description: description,
	}
}
