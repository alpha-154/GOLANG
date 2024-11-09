package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the world of Structs!")
	//no inheritance, no super class in golang

	johndoe := User{"John Doe", "johndoe@go.dev", 32, true}
	fmt.Println(johndoe)
	fmt.Printf("Type of johndoe is: %T \n", johndoe)
	fmt.Printf("User's profile : %+v \n", johndoe)
	fmt.Printf("User's name : %v and email : %v \n", johndoe.Name, johndoe.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	status bool
}
