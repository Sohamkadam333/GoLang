package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("First switch case program")

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
	case 4:
		fmt.Println("Number is 4")
		fallthrough  // fallthrough is use to go to next case without checking
	case 5:
		fmt.Println("Number is 5")
	case 6:
		fmt.Println("Number is 6")
	default:
		fmt.Println("Invalid Number")
	}
}