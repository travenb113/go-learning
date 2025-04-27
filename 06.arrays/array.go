package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8} // Defined lenghts
	var arr3 = [...]int{9, 10, 11}
	arr4 := [...]int{12, 13, 14, 15, 16} // Inferred lenghts

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
