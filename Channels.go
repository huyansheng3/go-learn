package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(time.Second )
		message <- "ping"
	}()


	fmt.Println("start")

	msg := <- message

	fmt.Println(msg)

	fmt.Println("end")

}