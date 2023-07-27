package readingfiles

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestPostsFromFS(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"hello-world.md": {Data: []byte(
				`Title: Hello, World!
Description: donkeys
Tags: tdd, go
---
Hello
body ody ody`)},
			//"hello-twitch.md": {Data: []byte("Title: Hello, Twitch!")},
		}
		posts, err := PostsFromFS(fileSystem)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fileSystem) {
			t.Errorf("Expected %d posts, got %d posts", len(fileSystem), len(posts))
		}

		assertPost(t, posts[0], Post{
			Title:       "Hello, World!",
			Description: "donkeys",
			Tags:        []string{"tdd", "go"},
			Body:        "Hello\nbody ody ody",
		})
	})

	t.Run("Error handling", func(t *testing.T) {
		_, err := PostsFromFS(FailingFS{})
		if err == nil {
			t.Error("expected error but didnt get one")
		}
	})
}

func assertPost(t *testing.T, got, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v but wanted %#v", got, want)
	}
}

type FailingFS struct{}

func (f FailingFS) Open(string) (fs.File, error) {
	return nil, errors.New("I always fail")
}
