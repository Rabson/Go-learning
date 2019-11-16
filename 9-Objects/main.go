package main

import (
	"fmt"
)

func main() {
	Post := map[string]interface{}{
		"ID":      1,
		"Title":   "Go Rocks!!",
		"name":    "yogesh",
		"surname": "nishad",
	}

	fmt.Println(Post)

	// to access values
	fmt.Println(Post["ID"])
	fmt.Println(Post["Title"])
	fmt.Println(Post["surname"])
	// ....
}

// output:
// map[ID:1 Title:Go Rocks!! name:yogesh surname:nishad]
// 1
// Go Rocks!!
// nishad
