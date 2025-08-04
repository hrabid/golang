package main

import (
	"fmt"
)
/* Naming Rules for go functions
* * Must start with a letter 
* * Only contain alpha numeric (A-Z, a-z, 0-9)
* * No spaces 
* * Case sensitive 
* * For multi word function name multi-word variable convention is recommended
*/

func myMessage()  {
	fmt.Println("Calling first function....\n")	
}

// function parameters & args 
func my_family(fname string, age int) {
	fmt.Println("Hello", fname, "Ahmad", ", Age:", age)
}


func main() {
	myMessage()
	my_family("Abid", 18)
	my_family("Afif",2)
	my_family("Hasan", 56)
}
