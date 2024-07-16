package main

import "fmt"

func main(){
	fmt.Println("First loop Code")

	days := []string{"Mon","Tue","Wed","Thu","Fri","Sat","Sun"}

	// for i := 0; i < len(days); i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index,value := range days{
		fmt.Printf("Index = %v, value = %v\n",index,value)
	}

}