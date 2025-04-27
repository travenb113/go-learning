package main

import "fmt"

func main() {
	fruits := [3]string{"apple", "banana", "mango"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
