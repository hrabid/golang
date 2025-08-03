package main

import (
	"fmt"
)

func main() {
	/* Ways to create a slice:
	* * Using the []datatype{values} format
	* * Create a slice from an array
	* * Using the make() function
	*/

	// Using the []datatype{values} format 
	myslice := []int{1,27,16}
	fmt.Println(myslice)
	fmt.Println("Lenth:",len(myslice))
	fmt.Println("Capacity:",cap(myslice))

	myslice1 := []int{}
	fmt.Println(myslice1)
	fmt.Println("Lenth:",len(myslice1))
	fmt.Println("Capacity:",cap(myslice1))

	myslice2 := []string{"Abid","Afif","Ahmad","Shafi"}
	fmt.Println(myslice2)
	fmt.Println("Lenth:",len(myslice2))
	fmt.Println("Capacity:",cap(myslice2))




}
