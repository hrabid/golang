package main

import (
	"fmt"
)

func main() {

	/* fmt.Print() function 
	* * Prints the args as it is 
	* * Can't add newline or space (manual)
	*/
	var name, age, session = "Afif", 18, 2025
	fmt.Print(name, "\n")
	fmt.Print(age, session, "\n")
	// Print() inserts a space if both var is a number 
	

	/* fmt.Println() function
	* * Can add newline & space by default 
	*/
	fmt.Println(name)
	fmt.Println(age, session)
	fmt.Println("\n")

	/* fmt.Printf() function
	* * Supports format specifiers 
	* * Can't add new line & space (manual)
	*/
	fmt.Printf("Name : %v, and type of var name is : %T" , name, name)
	fmt.Println("\n")

}
