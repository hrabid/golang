package main

import (
	"fmt"
)

// Return value is not necessary if you dont want to return any value 
func return_sum(x int, y int) int {
	return x + y 
}

// Nmaed return type 
func return_mul(a int, b int) (result int) {
	result = a * b
	return // or return result
}



func main() {
	fmt.Println(return_sum(6, 10))
	fmt.Println(return_mul(5,5))
	// storing return value in a variable
	total := return_mul(5,5)
	fmt.Println(total)
}
