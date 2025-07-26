package main

import "fmt"

func main() {
	fmt.Println("(=== Lets learn Method in Go Lang)")

	//No inheritence in golang; no parent no child
	//lets explore structure
	//Method are same as function like when we call function in classes
	//its become method same way in go lang we have struct so here in struct
	//we call them method .

	zahid := User{"Zahid", "zahidmiana56@gmail.com", true, 24}
	fmt.Println("Structure is : ", zahid)
	fmt.Printf("Zahid details are %+v\n", zahid)
	fmt.Printf("Name is: %v and Age is: %v\n", zahid.Name, zahid.Age)

	//method call
	zahid.GetStatus()
	zahid.NewMail()

	//here again calling check its change in orignal or just a copy
	fmt.Println("Structure is : ", zahid)

}

// struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// method
func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
