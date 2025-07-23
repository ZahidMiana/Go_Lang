package main

import "fmt"

func main() {
	fmt.Println("(===  Welcome to learn Maps in Go-Lang === )")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	languages["sol"] = "Solidity"

	fmt.Println("Our maps values are: ", languages)
	fmt.Println("Our maps values for Individuals: ", languages["JS"]) //wanna for single value

	delete(languages, "rb")
	fmt.Println("Our maps after deleting values are: ", languages)

	fmt.Println(".............Using Loop..................")
	// loops are interesting in go lang
	for key, value := range languages {
		fmt.Printf("For Key %v, Value is %v\n", key, value)
	}

}
