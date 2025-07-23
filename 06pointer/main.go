package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Pointer!")

	// var ptr *int
	// fmt.Println("Value of Pointer is, ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of Actual Pointer is: ", ptr)
	fmt.Println("Value of Actual Pointer is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is: ", myNumber)

}
