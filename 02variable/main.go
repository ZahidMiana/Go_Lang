package main

import "fmt"

const LoginToken string = "Hanji to kessy hein ap" //public

func main() {

	//string
	var username string = "Zahid Iqbal"
	fmt.Println(username)
	fmt.Printf("Variable is of Type %T \n", username)

	//bool
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	//int
	var smallVal uint16 = 266
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	//float
	var smallValfloat float32 = 250.3456
	fmt.Println(smallValfloat)
	fmt.Printf("Variable is of type: %T \n", smallValfloat)

	//default value and some alises
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type
	var website = "mianamart.dev"
	fmt.Println(website)

	//no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	//above public varibale print here
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
