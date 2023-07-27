package main

import (
	"fmt"
	"hello/readingfiles"
	"log"
	"os"
)

func main() {
	posts, err := readingfiles.PostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
}
