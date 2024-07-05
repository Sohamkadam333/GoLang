package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating of our Pizza: ")


	// ok , err := 
	// _ = ignore

	input,_ := reader.ReadString('\n')  // read till new line 'Enter'
	println("Thanks for rating our Pizza ",input)



}