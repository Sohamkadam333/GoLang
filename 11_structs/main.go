package main

import "fmt"

func main() {
	fmt.Println("My first struct lecture")

	soham := User{"Soham",23,false}
	fmt.Println(soham)
	fmt.Printf("Details = %+v\n",soham)

}

// use capital letters because we are going to use it outside.
type User struct{
	Name string
	Age int
	IsMarried bool
}
