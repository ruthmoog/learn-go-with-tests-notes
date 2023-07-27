package readingfiles

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
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
		Body:        readBody(scanner),
	}
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // skip line divider
	var buf bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
