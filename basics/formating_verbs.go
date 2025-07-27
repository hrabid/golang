package main

import (
	"fmt"
)

func main() {
	/* General formating verbs
	* * %v ঃ: prints the  value 
	* * %#v : prints the value in go-syntax 
	* * %T : 
	* * %% : prints % sign 
	*/
	var name string = "Ahmad"
	var marks int = 90
	math := 95.25
	fmt.Printf("Value of Variable: %v\n", name)
	fmt.Printf("Value of Variable in go-syntax: %#v\n", name)
	fmt.Printf("Type of the Value : %T\n", name)
	fmt.Printf("I got %v%% of makrks\n", marks)

	fmt.Printf("Type of math var : %T\n", math)
	fmt.Printf("\n")



	/* Integer formating verbs 
	* Verb	Description 
	* %b : Base 2(binary)
	* %d : Base 10 (decimal)
	* %+d : Base 10 and always show sign
	* %o : Base 8(octal)
	* %O : Base 8, with leading 0o
	* %x : Base 16(hex), lowercase
	* %X : Base 16, uppercase
	* %#x	Base 16, with leading 0x
	* %4d : Pad with spaces (width 4, right justified)
	* %-4d : Pad with spaces (width 4, left justified)
	* %04d : Pad with zeroes (width 4
	*/
	var i = 18
	fmt.Printf("Binary: %b\n", i)
	fmt.Printf("Decimal: %d\n", i)
	fmt.Printf("Decimal with sign : %+d\n", i)
	fmt.Printf("Ocatal: %o\n", i)
	fmt.Printf("Ocatal with leading 0o: %O\n", i)
	fmt.Printf("Hexadecimal: %x\n", i)
	fmt.Printf("Hexadecimal: %X\n", i)
	fmt.Printf("Hex with 0x: %#x\n", i)
	fmt.Printf("idk: %4d\n", i)
	fmt.Printf("idk: %-4d\n", i)
	fmt.Printf("Fills gap with zeros: %04d\n", i)
	fmt.Printf("\n")


	/* Floating point formating verbs
	* %e	Scientific notation with 'e' as exponent
	* %f : Decimal point, no exponent
	* %.2f : Default width, precision 2
	* %6.2f : Width 6, precision 2
	* %g : Exponent as needed, only necessary digits
	*/
	var x = 3.141526
	fmt.Printf("Scientific notation: %e\n", x)
	fmt.Printf("Decimal point: %f\n", x)
	fmt.Printf("Two digits after decimal: %.2f\n", x)
	fmt.Printf("Three digits after decimal: %.3f\n", x)
	fmt.Printf("idk: %6.2f\n", x)
	fmt.Printf("idk: %g\n", x)
	fmt.Printf("\n")



	/* String formating verbs
	* %s	Prints the value as plain string
	* %q	Prints the value as a double-quoted string
	* %8s	Prints the value as plain string (width 8, right justified)
	* %-8s	Prints the value as plain string (width 8, left justified)
	* %x	Prints the value as hex dump of byte values
	* % x	Prints the value as hex dump with spaces  
	*/
	var a = "Abdullah"
	fmt.Printf("Plain String: %s\n", a)
	fmt.Printf("Double-quoted string: %q\n", a)
	fmt.Printf("idk: %8s\n", a)
	fmt.Printf("idk: %-8s\n", a)
	fmt.Printf("Hexadecimal dump: %x\n", a)
	fmt.Printf("Hex dump with space: % x\n", a)
	fmt.Printf("\n")



	/* Boolewn formating verbs
	* %t : true or false (like %v)
	*/
	var m = true
	var n = false
	fmt.Printf("%t\n", m)
	fmt.Printf("%t\n", n)
	fmt.Printf("\n")


	
}
