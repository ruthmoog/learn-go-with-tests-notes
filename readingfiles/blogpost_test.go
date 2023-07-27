package readingfiles

import (
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hello, World!")},
		"hello-twitch.md": {Data: []byte("hello, Twitch!")},
	}
	posts := PostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("Expected %d posts, got %d posts", len(fs), len(posts))
	}
}
