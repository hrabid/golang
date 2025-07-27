package main

import (
	"fmt"
)

func main() {
	/* Basic Data types 
	* 1. Boolean
	* 2. Numeric 
	* 3. String
	*/

	// Boolean Data
	var b1 bool = true // typed, with initial value
	var b2 = true // untyped, with initial value
	var b3 bool // typed, without initial value
	b4 := true // untyped, with initial value

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	/* Integer Data
	* * Signed Integer : support Negative or positive 
	* * * int 
	* * * int8
	* * * int16
	* * * int32
	* * * int64 
	* * Unsigned Integer - non-nrgative values only 
	* * * uint
	* * * uint8
	* * * uint16 
	* * * uint32
	* * * uint64
	*/

	var a int = 15
	var b int = -4500 
	fmt.Printf("Type : %T, Value: %v\n", a,a)
	fmt.Printf("Type : %T, Value: %v\n", b,b)

	var c uint = 363737
	var d uint = 63738
	fmt.Printf("Type : %T, Value: %v\n", c,d)
	fmt.Printf("Type : %T, Value: %v\n", d,d)

	/* Float Data 
	* float32 
	* float64 
	*/
	e := 1.7e+308
	fmt.Printf("Type: %T, Value: %v\n", e, e)


	/* String Data 
	* string 
	*/
	var m string = "Hello"
	var n string
	var o string = "World"
	fmt.Printf("Type: %T, Value: %v\n", m, m)
	fmt.Printf("Type: %T, Value: %v\n", n, n)
	fmt.Printf("Type: %T, Value: %v\n", o, o)

}
