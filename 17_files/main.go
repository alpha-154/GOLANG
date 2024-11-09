package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang!!!")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Printf("The length of the file is: %v \n", length)
	defer file.Close()

	readFile("./myFile.txt")
}

func readFile(fileName string) {

	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data (in bytes) inside the file is: \n", databyte)
	fmt.Println("Text data (in string) inside the file is: \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
