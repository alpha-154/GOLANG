package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the maps tutorial!")
	// var colors map[string]string
	course := make(map[string]int)
	course["js"] = 1
	course["php"] = 2
	course["go"] = 3
	course["nextjs"] = 4
	fmt.Println(course)
	fmt.Println(course["js"])

	delete(course, "php")
	fmt.Println(course)

	for key, value := range course {
		fmt.Printf("%v : %v \n", key, value)
	}
}
