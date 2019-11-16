package main

import "fmt"

func main() {
	age := 11
	if age > 10 {
		fmt.Println("you are old")
	} else {
		fmt.Println("you are young")
	}

	gender := "female"
	switch gender {
	case "female":
		fmt.Println("your name is tonie ")
	case "male":
		fmt.Println("your name is tony")
	default:
		fmt.Println("this is default")
	}
}
