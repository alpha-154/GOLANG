package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the world of pointers in golang!")

	// var ptr *int
	// fmt.Println("Value of ptr is ", ptr) //initially ptr will be nil

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of number is ", myNumber)
	fmt.Println("Value of ptr is ", ptr)
	fmt.Println("Value of ptr is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Value of number is ", myNumber)
}
