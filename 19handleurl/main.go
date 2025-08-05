package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	fmt.Println("( ==== Hello lets learn how we handle url in go lang ==== )")

	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)

	fmt.Println("Scheme is: ", result.Scheme)
	fmt.Println("Host is: ", result.Host)
	fmt.Println("Path is: ", result.Path)
	fmt.Println("Port is: ", result.Port())
	fmt.Println("Raw Query is: ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("THe Type of Query params are: %T\n", qparams)
	//now
	fmt.Println(qparams["coursename"])

}
