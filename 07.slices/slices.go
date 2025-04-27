package main

import "fmt"

func main() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
	// The slice starts at the third element of the array, which has the value 12
	// (Remember that array indices start at 0. This means that [0] is the first element, [1] is the second element, and so on).
	// The slice can grow to the end of the array. This means that the capacity of the slice is 4.
}
