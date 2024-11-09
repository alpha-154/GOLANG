package main

import "fmt"

const LoginToken = "dadfude" //public variable since it's starts with capital letter.

func main() {

	var name string = "mrh rakib"
	fmt.Println(name)
	fmt.Printf("type of this variable is: %T \n", name)

	var isLoggedIn bool = true
	fmt.Printf("Variable is type of: %t \n", isLoggedIn)

	var smallVal int = 255
	fmt.Printf("The value of smallVall is %d \n ", smallVal)

	var smallFloat float32 = 255.9845
	fmt.Printf("The value of smallFloat is %.2f \n ", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "learngoprogramming.com"
	fmt.Println(website)

	//no var style
	numOfUsers := 56
	fmt.Println(numOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("LoginToken is of type: %T \n", LoginToken)
}
