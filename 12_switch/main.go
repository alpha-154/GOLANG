package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the switch in golang")
	seed := time.Now().UnixNano()
	fmt.Println("Value of seed is: ", seed)
	r := rand.New(rand.NewSource(seed))
	fmt.Println("Value of r is: ", r)
	diceNum := r.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1 and you can open 1 spot")
		fmt.Println("****")
	case 2:
		fmt.Println("Dice value is 2 and you can open 2 spot")
		fmt.Println("****")
	case 3:
		fmt.Println("Dice value is 3 and you can open 3 spot")
		fmt.Println("****")
	case 4:
		fmt.Println("Dice value is 4 and you can open 4 spot")
		fmt.Println("****")
	case 5:
		fmt.Println("Dice value is 5 and you can open 5 spot")
		fmt.Println("****")
	case 6:
		fmt.Println("Dice value is 6 and roll dice again")
		fmt.Println("****")
	default:
		fmt.Println("What was that?")
	}

}
