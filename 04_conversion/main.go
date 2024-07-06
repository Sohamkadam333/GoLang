package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter rating for pizza: ")

	input,_ := reader.ReadString('\n')

	fmt.Println("Thanks for rating")

	inputNumber,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Printf("Error converting num")
	} else {
		fmt.Println("New Rating = ",inputNumber+1)
	}


}