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
	
	/* len() function - returns the length of the slice (the number of elements in the slice)

	cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
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

	myslice2 := [7]string{"Abid","Afif","Ahmad","Shafi"}
	fmt.Println(myslice2)
	fmt.Println("Lenth:",len(myslice2))
	fmt.Println("Capacity:",cap(myslice2))

	// creating from existing array
	arr1 := [...]int{4,8,5,3,2,9,6,8}
	//myslice3 := arr1[1:3]
	myslice3 := arr1[1:6]
	fmt.Println(myslice3)
	fmt.Println("Lenth:", len(myslice3))
	fmt.Println("Capacity:", cap(myslice3))

	// creating slices by make() func 
	/* Syntax
	* slice_name := make([]int,length,capacity)
	*/
	myslice4 := make([]int,5,9)
	// myslice4 := make([]int,5)
	fmt.Println(myslice4)
	fmt.Println("Lenth:", len(myslice4))
	fmt.Println("Capacity:", cap(myslice4))




}
