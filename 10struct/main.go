package main

import "fmt"

func main() {
	fmt.Println("(=== Lets learn Struct in Go Lang)")

	//No inheritence in golang; no parent no child
	//lets explore structure

	zahid := User{"Zahid", "zahidmiana56@gmail.com", true, 24}
	fmt.Println("Structure is : ", zahid)
	fmt.Printf("Zahid details are %+v\n", zahid)
	fmt.Printf("Name is: %v and Age is: %v\n", zahid.Name, zahid.Age)

}

// struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
