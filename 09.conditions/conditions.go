package main

import "fmt"

func main() {
	// Example of use if, else and else if
	age := 19

	if age < 18 {
		fmt.Println("You aren't an adult.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a pensioner.")
	}

	// Example of use switch
	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Day off")
	}
}
