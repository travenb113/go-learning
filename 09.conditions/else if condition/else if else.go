package main

import "fmt"

func main() {
	time := 11
	// code to be executed if 10 is true
	if time < 10 {
		fmt.Println("Good morning.")
		// code to be executed if < 10 is false and < 20 is true
	} else if time < 20 {
		fmt.Println("Good day.")
		// code to be executed if < 10 and < 20 are both false
	} else {
		fmt.Println("Good evening.")
	}
}
