package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the go time package!")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
