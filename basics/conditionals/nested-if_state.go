package main

import (
	"fmt"
)

func main() {
	x := 11 
	if x >= 10 {
		fmt.Println("x is greater than 10....\n")
		if x >= 12 {
			fmt.Println("x is also greater than 12....\n")
		}
	} else {
		fmt.Println("x is less than 10...\n")
	}
}
