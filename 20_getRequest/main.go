package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to making GET request in golang!!!")

	//PerformGetRequest()

	//PerformPostJsonRequest()

	PerformPostFormRequest()
}

func PerformGetRequest() {
	serverUrl := "http://localhost:3000/get"
	response, err := http.Get(serverUrl)

	if err != nil {
		panic(err)
	}
	statusCode := response.StatusCode
	contentLength := response.ContentLength
	fmt.Println("Status code is: ", statusCode)
	fmt.Println("Content length is: ", contentLength)
	if statusCode != 200 {
		return
	}

	defer response.Body.Close()

	// databyte, _ := io.ReadAll(response.Body)
	// fmt.Println(string(databyte))

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	serverUrl := "http://localhost:3000/json"

	//fake json payload

	requestBody := strings.NewReader(`
   {
	"courseName": "Let's go with golang",
	"price": 0
   }	
   `)

	response, err := http.Post(serverUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	serverUrl := "http://localhost:3000/form"

	//formdata

	data := url.Values{}
	data.Add("courseName", "Let's go with golang formdata")
	data.Add("price", "0")
	data.Add("duration", "30")

	response, err := http.PostForm(serverUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
