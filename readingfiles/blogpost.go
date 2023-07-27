package readingfiles

import (
	"io/fs"
)

func PostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		posts = append(posts, makePostFromFile(fileSystem, f.Name()))
	}
	return posts, nil
}

func makePostFromFile(fileSystem fs.FS, fileName string) Post {
	blogFile, _ := fileSystem.Open(fileName)
	return newPost(blogFile)
}
