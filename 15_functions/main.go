package main

import "fmt"

func main() {
	fmt.Println("First Function code")

	result := adder(10,20)  // defination and initialization
	fmt.Printf("Result = %v\n",result)

	result = proAdder(10,20,30,40,50)
	fmt.Printf("Result = %v\n",result)
}

func adder(no1 int, no2 int) int{
	return no1+no2
}

func proAdder(values ...int) int {
total := 0

for i := 1; i < len(values);i++{
	total += values[i]
}
return total
}