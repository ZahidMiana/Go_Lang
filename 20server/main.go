package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("( ==== Now lets learn Creating Server in Go lang)")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World! Welcome to your first Go server.")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
