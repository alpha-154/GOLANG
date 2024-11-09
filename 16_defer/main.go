package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to defer in golang!!")

	fmt.Println("NO Defer") //maintains the LIFO order while working with defer keyword
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	deferFunction()

}

func deferFunction() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
