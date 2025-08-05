package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/ZahidMiana"

func main() {
	fmt.Println(" (=== Lets Learn Web Request In Golang === )")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() //Caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println("Content is: ", content)
}
