package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v & value is %v \n", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 6 {
			goto lco
		}
		if rougueValue == 5 {
			break
		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumping at lco")

}
