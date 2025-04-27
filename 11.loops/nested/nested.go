package main

import "fmt"

func main() {
	abc := [2]string{"big", "tasty"}
	def := [5]string{"apple", "banana", "orange", "pineapple", "mango"}
	for i := 0; i < len(abc); i++ {
		for u := 0; u < len(def); u++ {
			fmt.Println(abc[i], def[u])
		}
	}
}
