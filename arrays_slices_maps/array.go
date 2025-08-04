package main

import (
	"fmt"
)

func main() {
	/* Syntax of array 
	* var arr_name = [lenth]datatype{values}
	* arr_name := [lenth]datatype{values}
	* // lenth is defuned 
	* var arr_name = [...]datatype{values} 
	* arr_name := [...]datatype{values}
	* // lenth is inferred 
	*/

	var arr1 = [3]int{12,15,17}
	arr2 := [2]int{25,27}
	// lenth is defined 
	fmt.Println(arr1)
	fmt.Println(arr2)

	var arr3 = [...]int{5,8,9,4,3}
	arr4 := [...]int{2,28,273,2738,73}
	// compilet decides the lenth of array 
	fmt.Println(arr3)
	fmt.Println(arr4)

	var arr5 = [...]string{"Abid","Afif","Ahmad","Umayer"}
	// arr5 := [...]string{"Abid","Afif","Ahmad","Umayer"}
	fmt.Println(arr5)

	// Accessing array values
	students := [...]string{"Asif","Nahid","Hasnat","Quader","Shadik","Adnan"}
	fmt.Println(students[0])
	fmt.Println(students[2])

	// Changing Elements of array
	fmt.Println("Before Modification:", students[3])
	students[3] = "Abdul Quader"

	// Array initialization 
	arr6 := [5]int{} // not initialized 
	fmt.Println(arr6)
	arr7 := [5]int{5,7,8} // partially initialized
	fmt.Println(arr7)
	arr8 := [5]int{5,7,8,9,4} // fully initialized 
	fmt.Println(arr8)

	// Initializing Specific element of array
	var arr9 = [...]string{1:"Hasib",10:"Hasin"}
	fmt.Println(arr9)
	var arr10 = [5]int{1:12,4:25}
	fmt.Println(arr10)

	// lenth of array
	fmt.Println("The len() func show the number of elemets in array")
	fmt.Println(len(arr1))
	fmt.Println(len(arr10))



}
