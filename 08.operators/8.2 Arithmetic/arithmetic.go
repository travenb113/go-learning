package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 2
	x++            // Increments x by 1
	fmt.Println(x) // Prints the value of x
	x--            // Decrements x by 1
	fmt.Println(x) // Prints the value of x after decrementation

	fmt.Println(x + y) // Operator +. Name: Addition, Adds together two values
	fmt.Println(x - y) // Operator -. Name: Subtraction, Subtracts one value from another
	fmt.Println(x * y) // Operator *. Name: Multiplication, Multiplies two values
	fmt.Println(x / y) // Operator /. Name: Division, Divides one value by another
	fmt.Println(x % y) // Operator %. Name: Modulus, Returns the division remainder
}
