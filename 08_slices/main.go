package main

import (
	"fmt"
	"sort"
)

// The main function demonstrates the use of slices in Go programming language.
// It begins by declaring a slice of strings for fruits and then appending new
// fruits to the slice. It also demonstrates how to append a slice to another
// slice and how to use the make function to create a slice with a specified
// size. Additionally, the code shows how to sort a slice of integers in
// ascending order and how to check if a slice is sorted using the
// sort.IntsAreSorted method.
func main() {
	fmt.Println("Welcome to the slices tutorial!")

	var fruitList = []string{"apple", "tomato", "potato"}

	fmt.Printf("Type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "mango")
	fruitList = append(fruitList, "banana")
	fmt.Println("List of fruits: ", fruitList)
	fmt.Println("List of fruits: ", len(fruitList))

	fruitList = append(fruitList[1:3])
	fmt.Println("List of fruits: ", fruitList)
	fmt.Println("List of fruits: ", len(fruitList))

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 123
	highScores[3] = 567
	//highScores[4] = 678
	highScores = append(highScores, 989, 812)

	fmt.Println(highScores)
	fmt.Println(len(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove values from slices based on index
	var courses = []string{"reactjs", "angular", "vuejs", "flutter"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	fmt.Println(len(courses))
}
