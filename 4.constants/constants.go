package main

import "fmt"

const (
	A int = 1     // A is a typed integer constant with the value 1
	B     = 3.14  // B is an untyped float constant with the value 3.14
	C     = "Hi!" // C is an untyped string constant with the greeting "Hi!"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
