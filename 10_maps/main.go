package main

import (
	"fmt"
)

func main() {
	fmt.Println("First Map lecture")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["C"] = "C lang"

	fmt.Println("Values in map = ",languages)
	fmt.Println("Value of JS key is ", languages["JS"])

	// delete key value
	delete(languages,"JS")
	fmt.Println("Values in map after deleting = ",languages)


}