package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the world of if else in go!")

	loginCount := 23
	var result string

	if loginCount > 15 {
		result = "loginCount is more than 15"
	} else {
		result = "loginCount is less than 15"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else if num < 20 {
		fmt.Println("Number is less than 20")
	} else {
		fmt.Println("Number is less or greater than 20")
	}
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
