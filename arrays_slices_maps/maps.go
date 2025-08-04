package main

import (
	"fmt"
)

func main() {
	/* Syntax 
	* var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	* b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	*/
	//var a = map[string]string{"name": "Abid","age":"19","job":"Assistant"}
	//b := map[string]string{"brand":"ford","model":"mustang","year":"1994"}
	
	// make() function for creating map 
	var a = make(map[string]string) // creates an empty map
	a["name"] = "Abid"
	a["age"] = "19"
	a["Education"] = ""

	b := make(map[string]string)
	b["brand"] = "ford"
	b["model"] = "mustang"

	// Print whole map 
	fmt.Printf("map a : %v\n",a)
	fmt.Printf("map b : %v\n",b)

	// Access elements of map  
	fmt.Println(a["name"])

	// updating value 
	a["name"] = "Ahmad"
	fmt.Println(a["name"])

	// Adding an element 
	a["Job"] = "Engineer"
	fmt.Println(a["Job"])
	fmt.Println(a)

	// delete an element 
	delete(a,"Job")
	fmt.Println(a)


	// Checking for existance of a key & value 
	val1, ok1 := a["name"]
	val2, ok2 := a["age"]
	val3, ok3 := a["Job"]
	_, ok4 := a["Education"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3) // these three are checking for key & value
	fmt.Println(ok1)
	fmt.Println(ok4) // only checking for key not value
	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println("\n") // only checking for value not key 

	

	// defference between make() function & without make() function
	fmt.Println("Difference between using make() function & not using it.....")
	var c = make(map[string]int)
	var d map[string]int
	fmt.Println(c == nil)
	fmt.Println(d == nil)

	/* Supported key types 
	* Booleans
	* Numbers
	* Strings
	* Arrays
	* Pointers
	* Structs
	* Interfaces (as long as the dynamic type supports equality)
	*/

	/* Invalid key types 
	* Slices 
	* Maps 
	* Functions 
	*/

	/* Maps are references to hash tables.
	* If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	*/
	e := map[string]int{"hasib":15,"habib":17,"afif":5,"mihhaz":19}
	f := e
	fmt.Println(e)
	f["hasib"] = 22
	fmt.Println("After changes in variable f: ")
	fmt.Println(e)

	// iterating over maps 
	for k, v := range f {
		fmt.Printf("%v %v\n", k, v)
	}



}
