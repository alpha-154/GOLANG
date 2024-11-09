package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/2"

func main() {
	fmt.Println("Welcome to the handling web requests in golang!!!")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T \n", response)

	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", response.StatusCode, response.Status))
	}

	databyte, err := io.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databyte)
	fmt.Printf("Content is of type: %T \n", content)
	fmt.Printf("Content is: %v \n", content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
