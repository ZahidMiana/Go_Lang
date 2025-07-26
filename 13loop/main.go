package main

import "fmt"

func main() {
	fmt.Println("(=== Lets learn Loop ==== )")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	//for loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//other way
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//its like a for eaach loop
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco //we create below so when met its move to our goto statement
		}

		if rougueValue == 5 {
			rougueValue++
			//continue (its just skip the value and iterate next)
			break //its break the loop when condition met
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	// its basically goto statement like when condition met we trigger our goto statement
lco:
	fmt.Println("Jumping at LearnCodeonline.in")
}
