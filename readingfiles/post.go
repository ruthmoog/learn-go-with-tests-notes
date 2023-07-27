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

	readline := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	return Post{
		Title:       readline("Title: "),
		Description: readline("Description: "),
	}
}
