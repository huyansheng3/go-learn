package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string, 2)

	go func() {
		time.Sleep(time.Second)
		message <- "111"
		time.Sleep(time.Second)
		message <- "222"

	}()


	fmt.Println("start1")

	fmt.Println(<-message)

	fmt.Println("end1")

	fmt.Println("start2")

	fmt.Println(<-message)

	fmt.Println("end2")
}
