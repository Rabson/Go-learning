package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// key value pairs
	kvs := map[string]string{
		"name":    "yogesh",
		"surname": "nishad",
	}

	for key, value := range kvs {
		fmt.Println(key, value)
	}
}

// output:
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// name yogesh
// surname nishad
