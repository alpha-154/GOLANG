package main

import (
	"fmt"
)

func main() {
	var arr [3]string

	arr[0] = "Hello"
	arr[1] = "World"
	arr[2] = "!"

	fmt.Println(arr[0], arr[1], arr[2])
	fmt.Println(arr)
	fmt.Println(len(arr))

	var vegList = [3]string{"Potato", "Carrot", "Beans"}
	fmt.Println("vegy list is: ", vegList)
}
