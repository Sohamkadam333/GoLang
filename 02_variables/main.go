package main

import "fmt"

// globale variable cannot use walrus operator
// JwtToken := "skfjq42ef5wefuo"

var jwtToken string = "skfjq42ef5wefuo"

// if variables first letter is capital then it is public variable.

const AuthToken string = "hgnr54t35g4a53w";   // public variable

func main() {
	var username string = "SohamK"
	fmt.Println(username)
	fmt.Printf("Type of username variable is %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn variable is %T \n",isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Type of smallInt variable is %T \n",smallInt)

	var smallFloat float32 = 255.5454545454
	fmt.Println(smallFloat)
	fmt.Printf("Type of smallFloat variable is %T \n",smallFloat)

	// default values and aliases
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Type of defaultInt variable is %T \n",defaultInt)

	// implicit type
	var fName = "Soham"
	fmt.Println(fName)
	fmt.Printf("Type of fName variable is %T \n",fName)

	// no var style

	lName := "Kadam"
	fmt.Println(lName)
	fmt.Printf("Type of lName variable is %T \n",lName)


	fmt.Print("JwtToke = ",jwtToken)



}