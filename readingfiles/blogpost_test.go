package readingfiles

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello-world.md": {Data: []byte("Title: Hello, World!")},
			//"hello-twitch.md": {Data: []byte("Title: Hello, Twitch!")},
		}
		posts, err := PostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("Expected %d posts, got %d posts", len(fs), len(posts))
		}

		expectedFirstPost := Post{Title: "Hello, World!"}
		if posts[0] != expectedFirstPost {
			t.Errorf("expected %#v, but got %#v", expectedFirstPost, posts[0])
		}
	})

	t.Run("Error handling", func(t *testing.T) {
		_, err := PostsFromFS(FailingFS{})
		if err == nil {
			t.Error("expected error but didnt get one")
		}
	})
}

type FailingFS struct{}

func (f FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("I always fail")
}
