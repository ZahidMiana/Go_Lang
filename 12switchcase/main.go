package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("(====  Switch Statement in Go Lang)")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can Open")
	case 2:
		fmt.Println("Dice Value is 2 and you can spot")
	case 3:
		fmt.Println("Dice Value is 3 and you can SPot")
	case 4:
		fmt.Println("Dice Value is 4 and this is pretty good")

	case 5:
		fmt.Println("Dice Value is 5 and this is amazing")
	case 6:
		fmt.Println("Dice Value is 6 and Chakka Hurrah!")
	default:
		fmt.Println("WHat was that!")
	}

}
