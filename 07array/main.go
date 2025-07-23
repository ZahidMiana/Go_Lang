package main

import "fmt"

func main() {
	fmt.Println("(=== Welcome to Array in Go Lang===)")

	var myArray = [8]int{12, 14, 16, 18, 20}

	myArray[5] = 22
	myArray[6] = 26

	fmt.Println("My Array is: ", myArray)
	fmt.Println("Length of Array is ", len(myArray))

	vegetables := [4]string{"potato", "beans", "kella", "kaddo"}
	fmt.Println("Vegatables are: ", vegetables)
}
