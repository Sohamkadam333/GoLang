package main

import "fmt"

func main(){
	fmt.Println("First if else program")

	age := 12

	if age >= 18{
		fmt.Println("You can drive")
	} else {
		fmt.Println("You cannot Drive")
	}

	if num := 3; num > 2 {
		fmt.Println("Number is greater than 2")
	}

	// err := 22

	// if err != nil {
	// 	fmt.Println("Error is not nil")
	// }

}