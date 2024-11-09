package main

import (
	"fmt"
	"net/url"
)

const testUrl string = "https://localhost:3000/courses?category=physics&duration=long"

func main() {

	fmt.Println("Welcome to handling url in golang!!!")
	fmt.Println(testUrl)

	//parsing
	result, _ := url.Parse(testUrl)

	fmt.Println("Schema: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of qparams is: %T \n", qparams)
	fmt.Println(qparams["category"])
	fmt.Println(qparams["duration"])

	for _, val := range qparams {
		fmt.Println("Current param is: ", val)
	}

	constructUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost:3000",
		Path:     "/courses",
		RawQuery: "category=physics&duration=long",
	}

	fmt.Println("constructed url: ", constructUrl.String())

}
