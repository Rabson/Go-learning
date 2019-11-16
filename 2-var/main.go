package main

import "fmt"

const num = 3 // declaring a constant

func main() {
	var flt float64 = 2.34     // declaring a float
	var name string = "Scotch" // declaring a string
	a, b := true, false        // shorthand declaration syntax
	fmt.Println(flt)
	fmt.Println(name)
	fmt.Println(a, b)

}

// output:
// 2.34
// Scotch
// true false
