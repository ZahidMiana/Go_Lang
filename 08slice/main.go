package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("(==== Welcome to SLices ==== )")

	var fruit = []string{"kella", "mango", "amrood"}

	fmt.Println("Our Slice is: ", fruit)
	fmt.Printf("Type of Fruit is:  %T\n ", fruit)

	//add values
	fruit = append(fruit, "Jaman", "nashpatti")
	fmt.Println("Now After Adding Values: ", fruit)

	//other way for adding
	fruit = append(fruit[1:3])
	fmt.Println("Now After slicing Values: ", fruit)

	//use make keyword for making slice
	highScore := make([]int, 4)

	highScore[0] = 230
	highScore[1] = 260
	highScore[2] = 290
	highScore[3] = 300
	fmt.Println("HY this slice is using make method: ", highScore)
	fmt.Printf("Type of highScore: %T\n", highScore)

	highScore = append(highScore, 555, 666, 777)
	fmt.Println("HY this slice is using append method: ", highScore)

	//sort method in slice
	sort.Ints(highScore)
	fmt.Println("Values after sorted: ", highScore)
	fmt.Println("It gave true because we sort the values:", sort.IntsAreSorted(highScore))

	//how to remove values from slice
	var courses = []string{"python", "ruby", "c++", "java", "js", "go lang"}
	fmt.Println("Courses are : ", courses)

	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses are : ", courses)

}
