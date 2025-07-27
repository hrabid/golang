package main

import (
	"fmt"
)

func main() {
	// Constants are read only 
	const x int = 12
	fmt.Println(x)
	/* Note: there is no way to initialize a Constant without a value 
	*/

	const PI = 3.141526
	fmt.Println(PI)

	/* Constant Rules 
	* * Variable like naming schme 
	* * Upper case is best paractice for differenciation
	* * Can be declared inside or out side of a func
	*/

	/* Types of constants 
	* 1. Typed 
	* 2. Untyped 
	*/
	const NAME string = "Hasan"
	const SKILL = "Golang"
	fmt.Println(NAME)
	fmt.Println(SKILL)

	// Multiple Constant Declaration 
	const (
		A int = 17
		B string = "Hello"
		C bool = true 
		D = 3.45
	)
	fmt.Println(A, B, C, D)

}
