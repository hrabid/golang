package main

import (
	"fmt"
)
func main() {
	x := 246
	fmt.Printf("x = %v\n Decimal: %d\n Binary: %b\n Octal: %o\n Hex: %x,\n",x, x,x,x,x)

	// hex 
	greet := "Hello, World!"
	fmt.Printf("greet: %x\n",greet)
	for i := 0; i < len(greet); i++ {
		fmt.Printf("%x",greet[i])
	}
	fmt.Printf("\n")
}
