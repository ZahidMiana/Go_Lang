package main

import "fmt"

func main() {
	fmt.Println("( ==== Lets Learn Defer ==== )")

	fmt.Println("Start")
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	defer fmt.Println("Defer 4")
	fmt.Println("End")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
