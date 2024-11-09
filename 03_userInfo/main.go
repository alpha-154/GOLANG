package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user Input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || comma error
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", input)
	fmt.Printf("Thanks for rating us %T", input)
}
