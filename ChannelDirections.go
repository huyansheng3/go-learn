package main

import "fmt"

func ping(pings chan<- string, message string )  {
	pings <- message
}

func pong(pings <-chan  string, pongs chan<- string)  {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)

	pongs := make(chan string, 1)

	ping(pings, "hello world")

	pong(pings, pongs)

	fmt.Println(<-pongs, pongs, pings)



}