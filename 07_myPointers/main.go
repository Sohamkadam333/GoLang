package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Print("Default ptr value = ",ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of ptr = ",ptr)
	fmt.Println("Value of *ptr = ",*ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of *ptr = ",*ptr)

}