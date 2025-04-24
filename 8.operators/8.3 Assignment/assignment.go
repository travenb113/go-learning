package main

import "fmt"

func main() {
	x := 10 // In double form: 1010
	y := 2  // In double form: 0010

	// Simple assignment example
	fmt.Println("Изначальное значение x:", x)

	// Operator +=
	x += y // x = x + y
	fmt.Println("После x += y, значение x:", x)

	// Operator -=
	x -= y // x = x - y
	fmt.Println("После x -= y, значение x:", x)

	// Operator *=
	x *= y // x = x * y
	fmt.Println("После x *= y, значение x:", x)

	// Operator /=
	x /= y // x = x / y
	fmt.Println("После x /= y, значение x:", x)

	// Operator %=
	x %= y // x = x % y
	fmt.Println("После x %= y, значение x:", x)

	// Bitwise reversal operators
	x = 10 // Reset x for new operations
	fmt.Println("Изначальное значение x для битовых операций:", x)

	// Operator |=
	x |= y                                      // x = x | y
	fmt.Println("После x |= y, значение x:", x) // 1010 | 0010 = 1010 (10)

	// Operator ^=
	x ^= y                                      // x = x ^ y
	fmt.Println("После x ^= y, значение x:", x) // 1010 ^ 0010 = 1000 (8)

	// Operator <<=
	x <<= 1                                      // x = x << 1 (shift left by 1)
	fmt.Println("После x <<= 1, значение x:", x) // 1000 << 1 = 10000 (16)

	// Operator >>=
	x >>= 1                                      // x = x >> 1 (shift right by 1)
	fmt.Println("После x >>= 1, значение x:", x) // 10000 >> 1 = 1000 (8)

	// Increment and decrement
	x++ // Increment x
	fmt.Println("После инкремента, значение x:", x)

	x-- // Decrement x
	fmt.Println("После декремента, значение x:", x)
}
