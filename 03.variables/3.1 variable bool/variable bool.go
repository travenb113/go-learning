package main

import "fmt"

func main() {
	var b1 bool = true // Typed declaration with initial value
	var b2 = true      // Untyped declaration with initial value
	var b3 bool        // Typed declaration without initial value
	b4 := true         // Untyped declaration with initial value

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
}
