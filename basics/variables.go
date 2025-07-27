package main 

import (
	"fmt"
)

func main()  {

	// syntax: two types of variable declaration 
	// 1. using var keyword
	// var var_name type = value
	var name string = "Abid"
	fmt.Println(name)
	// 2. using := sign 
	// var_name := value 
	age := 18
	fmt.Println(age)

	// variable initialization (without a value)
	var a string
	var b int 
	var c bool 

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// value assignment after initialization
	a = "Golang string"
	fmt.Println(a)


// difference between var & :=

/* var Can be used inside and outside of functions
* Variable declaration and value assignment can be done separately
*/

/* := Can only be used inside functions
* Variable declaration and value assignment cannot be done separately (must be done in the same line) 
*/

// Multiple Variable declaration
var m, n, o int = 12, 17, 18
x, y, z := "Abid", 28, true
fmt.Println(m, n, o)
fmt.Println(x, y, z)
/* Note: it is not possible to declare multiple types of varible after typing the type like int, bool etc
* if type is inferred then it is possible 
* also var keyword & := cannot be used in same varible
*/

// Variable declaration in a block 
var (
	e float32
	f int = 35
	g string = "Afif"
	h bool = true
)
fmt.Println(e, f, g, h)


/* Variable Naming Rules
* * start with a letter or (_)
* * can't start with a digit 
* * Alpha numerica only (a-z, A-Z, 0-9, _)
* * Variable Names are case sensitive 
* * No space 
* * can't be go keyword
*/
// Camel Case 
myVariableName := "Afif"
// Pascal Case
MyVariableName := "Afif"
// Snake Case 
my_variavle_name := "Afif"
fmt.Println(myVariableName, MyVariableName, my_variavle_name)
}
