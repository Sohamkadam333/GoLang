package main

import "fmt"

func main() {

	fmt.Println("First Array lecture")

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "mango"
	fruits[2] = "kiwi"

	fmt.Println("Fruit list = ",fruits)
	fmt.Println("Length of fruits = ",len(fruits))

	// declaraing and assigning values
	var games = [2]string{"Football","Cricket"}
	fmt.Println("Value in games list = ", games)

}
