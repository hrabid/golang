package main

import (
	"fmt"
)

func main() {
	/* go supports :
	* Comparison Operators: ==, !=, <. >, <=, >=
	* Logical Operators: &&, ||, ! 
	*/

	/* Conditional statements 
	* if statements
	* if else statements
	* else if statements
	* switch statements
	*/

	// the result of a conditional can be either true or false 
	x := 5
	y := 7 
	z := false 
	fmt.Println(x > y)
	fmt.Println(x != y)
	fmt.Println((x < y) && (x > y))
	fmt.Println((x < y) || z)
}
