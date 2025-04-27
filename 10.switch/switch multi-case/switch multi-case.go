package main

import "fmt"

func main() {
	day1 := 8

	switch day1 {
	case 1, 2, 3:
		fmt.Println("Monday", "Thursday", "Wednseday")
	case 4, 5, 6, 7:
		fmt.Println("Thursday", "Fridsay", "Saturday", "Sunday")
	default:
		fmt.Println("Not a weekday!")
	}
}
