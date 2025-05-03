package main

import "fmt"

func main() {
	asd := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "banana", "orange"}
	for i := 0; i < len(asd); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(asd[i], fruits[j])
		}
	}
}
