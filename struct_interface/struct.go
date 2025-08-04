package main

import (
	"fmt"
)
type Person struct {
		name string 
		age int 
		job string
		salary int
	}
func main() {
	// intitializing variable with Person type 
	var per1 Person
	var per2 Person

	// per1 specs 
	per1.name = "Abid"
	per1.age = 18
	per1.job = "Engineer"
	per1.salary = 8000

	// per2 specs
	per2.name = "Afif"
	per2.age = 12
	per2.job = "Student"
	per2.salary = 500

	fmt.Println("Name:", per1.name)
	fmt.Println("age:", per1.age)
	fmt.Println("Job:", per1.job)
	fmt.Println("Salary:", per1.salary)

	fmt.Println("\n")

	fmt.Println("Name:", per2.name)
	fmt.Println("age:", per2.age)
	fmt.Println("Job:", per2.job)
	fmt.Println("Salary:", per2.salary)
	fmt.Println("\n")

	// struct as function argument 
	printPerson(per1)
	fmt.Println("\n")
	printPerson(per2)
}

func printPerson(pers Person) {
	fmt.Println("Name:", pers.name)
	fmt.Println("age:", pers.age)
	fmt.Println("Job:", pers.job)
	fmt.Println("Salary:", pers.salary)
}
