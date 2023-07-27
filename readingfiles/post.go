package readingfiles

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	readline := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	return Post{
		Title:       readline(titlePrefix),
		Description: readline(descriptionPrefix),
		Tags:        strings.Split(readline(tagsPrefix), ", "),
	}
}
