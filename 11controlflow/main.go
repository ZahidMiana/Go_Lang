package main

import "fmt"

func main() {
	fmt.Println("(=== Lets learn if else statement in Go Lang ===)")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	//even odd
	if 9%2 == 0 {
		fmt.Println("Number is even ")
	} else {
		fmt.Println("Number is odd")
	}

	//other syntex
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

}
