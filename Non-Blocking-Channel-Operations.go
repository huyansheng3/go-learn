package main

import (
	"fmt"
)

func main() {
	messages :=make(chan  string)
	//signals :=make(chan string)


	select {
	case msg:= <- messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}


	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}





}