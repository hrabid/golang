package main

import (
	"fmt"
)

func main() {
	// Single case switch statement 
	day := 0
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Not a week day")
	}

	// multi case switch statement
	char := "l"
	switch char {
	case "a","e","i","o","u":
		fmt.Printf("The character %v is a vowel...\n", char)
	default:
		fmt.Printf("The character %v is a Consonant...\n", char)
	}

}
