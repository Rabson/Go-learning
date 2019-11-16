package main

import (
	"fmt"
)

type Post struct {
	ID      int
	Title   string
	name    string
	surname string
}

func main() {
	// create an instance of the Post
	p := Post{
		ID:      1,
		Title:   "Go Rocks!!",
		name:    "yogesh",
		surname: "nishad",
	}

	fmt.Println(p)

	// to access values
	fmt.Println(p.ID)
	fmt.Println(p.Title)
	fmt.Println(p.name)
	fmt.Println(p.surname)
	// ....
}

// output:
// {1 Go Rocks!! yogesh nishad}
// 1
// Go Rocks!!
// yogesh
// nishad
