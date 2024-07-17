package main

import "fmt"

func main() {

	u1 := User{23, "Soham"}
	fmt.Printf("Name = %v, age = %v\n",u1.Name,u1.Age)
	u1.UpdateAge(12)
	u1.GetAge();

}

type User struct {
	Age  int
	Name string
}

func (u User) GetAge() {
	fmt.Printf("Age is %v\n",u.Age)
}

func (u User) UpdateAge(newAge int) {
	u.Age = newAge
}