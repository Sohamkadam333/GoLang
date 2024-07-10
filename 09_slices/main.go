package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("First Slice lecture")

	// Syntax of slices is similar to arrays
	// But we do not mention how many value are come in
	// When you define value its array
	// slice is similar to dynamic size array

	var fruits = []string{"apple","banana"}
	fmt.Println("Fruits = ",fruits)

	fruits = append(fruits, "Kiwi","Papaya")
	fmt.Println("Fruits = ",fruits)

	fruits = append(fruits[1:2])
	fmt.Println("Fruits = ",fruits)

	// dynamic memory allocation

	highSocores := make([]int, 4)

	highSocores[0] = 561
	highSocores[1] = 6361
	highSocores[2] = 26323
	highSocores[3] = 4112

	fmt.Println("Value in highScore ",highSocores)

	// Allocate new memory

	highSocores = append(highSocores, 23,456)
	fmt.Println("Value in highScore ",highSocores)

	sort.Ints(highSocores)
	fmt.Println("Sorted value = ",highSocores)


	// Remove values based on the index

	var numbers = []int{1,2,3,4,5}
	fmt.Println("Values in number is ",numbers)

	index := 2
	numbers = append(numbers[:index], numbers[index+1:]... )
	fmt.Println("Values in number is ",numbers)
}