package main

import "fmt"

func main() {
	var num1 float32 = 3.1415926535
	fmt.Println("float32:", num1) // Outputs value with lower precision

	var num2 float64 = 3.1415926535
	fmt.Println("float64:", num2) // Outputs value with higher precision
}
