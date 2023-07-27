package readingfiles

import (
	"io/fs"
)

func PostsFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for _, f := range dir {
		posts = append(posts, makePostFromFile(fileSystem, f.Name()))
	}
	return posts
}

func makePostFromFile(fileSystem fs.FS, fileName string) Post {
	blogFile, _ := fileSystem.Open(fileName)
	return newPost(blogFile)
}
