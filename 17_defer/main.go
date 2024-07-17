package main

import "fmt"

func main(){

	// Defer Works in LIFO format 
	defer fmt.Println("First Defer Program") // 1

	defer fmt.Println("One\n") // 2
	defer fmt.Println("Two\n") // 3
	defer fmt.Println("Three\n") // 4

	fmt.Println("Hello World\n")

	PrintOneTo10()
}

func PrintOneTo10(){
	for i := 0;i<10;i++{
		defer fmt.Println(i) // 5
	}
}



