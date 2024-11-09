package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to functions in golang!")

	greet()

	result1 := add(2, 3)
	fmt.Println("Result is: ", result1)

	result2 := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Pro Adder result is: ", result2)

	result3, msg := twoArgs(1, 2)
	fmt.Println("Result 3 is: ", result3)
	fmt.Println("Message is: ", msg)

}

func greet() {

	fmt.Println("Hello, World!")
}

func add(a int, b int) int {

	return a + b
}

func proAdder(values ...int) int {

	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func twoArgs(x int, y int) (int, string) {

	return x + y, "Hello"
}
