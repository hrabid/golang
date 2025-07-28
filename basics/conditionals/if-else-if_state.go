package main

import (
	"fmt"
)

func main() {
	class := 14
	if class == 10 {
		fmt.Println("Hasib is ssc examinee")
	} else if class == 8 {
		fmt.Println("Hsib is a jsc examinee")
	} else if class == 12 {
		fmt.Println("Hasib is hsc examinee")
	} else if class == 16 {
		fmt.Println("Hasib is honors examinee")
	} else {
		fmt.Println("Hasib is masters examinee")
	}
}
