package main

import (
	"fmt"
)

func main() {
	/* Arithmetic Operators
	* * Addition : +
	* * Substraction : -
	* * Multiplication : *
	* * Division : /
	* * Modulus : % 
	* * Increment : ++
	* * Decrement : --
	*/
	x := 25
	y := 20
	add := x + y 
	sub := x - y 
	mul := x * y
	div := x / y
	mod := x % y 
	x++
	y--
	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
	fmt.Println(mod)
	fmt.Println(x)
	fmt.Println(y)
}
