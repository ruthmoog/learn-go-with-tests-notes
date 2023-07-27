package readingfiles

import (
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte("Title: Hello, World!")},
		//"hello-twitch.md": {Data: []byte("Title: Hello, Twitch!")},
	}
	posts := PostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("Expected %d posts, got %d posts", len(fs), len(posts))
	}

	expectedFirstPost := Post{Title: "Hello, World!"}
	if posts[0] != expectedFirstPost {
		t.Errorf("expected %#v, but got %#v", expectedFirstPost, posts[0])
	}
}
