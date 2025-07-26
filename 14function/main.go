package main

import "fmt"

func main() {
	fmt.Println(" ==== (Lets Learn Function in Go Lang === )")

	//function call
	greeter()
	greeter2()

	result := adder(10, 30)
	fmt.Println("Result of Addition is: ", result)

	proResult := ProAdder(2, 5, 6, 7, 10, 12)
	fmt.Println("ProAdder result is: ", proResult)

}

func greeter2() {
	fmt.Println("Hanji Bhai mein ty theek tussi sunao")
}

func greeter() {
	fmt.Println("O ki haal hai Go Lang")
}

func adder(a int, b int) int {
	return a + b
}

func ProAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}

	return total
}
