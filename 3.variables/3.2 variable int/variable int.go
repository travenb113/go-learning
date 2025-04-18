package main

import "fmt"

func main() {
	var x int = 500
	var y int = -4500 // Signed integers
	var x1 int = 500
	var y1 int = 4500 // Unsigned integers
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)
	fmt.Printf("Type: %T, value: %v\n", x1, x1)
	fmt.Printf("Type: %T, value: %v\n", y1, y1)
}
