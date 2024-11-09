package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the world of Methods!")
	//no inheritance, no super class in golang

	johndoe := User{"John Doe", "johndoe@go.dev", 32, true}
	fmt.Println(johndoe)
	fmt.Printf("Type of johndoe is: %T \n", johndoe)
	fmt.Printf("User's profile : %+v \n", johndoe)
	fmt.Printf("User's name : %v and email : %v \n", johndoe.Name, johndoe.Email)

	johndoe.GetStatus()
	johndoe.NewMail()

	fmt.Printf("User'semail : %v \n", johndoe.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "newmail@go.dev"
	fmt.Println("Is user active: ", u.Email)
}
