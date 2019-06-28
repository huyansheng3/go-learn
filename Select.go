package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "c1"
	}()


	go func() {
		time.Sleep(time.Second *3)
		c2 <- "c2"
	}()


	now := time.Now()

	for i:=0; i<2 ; i++  {
		fmt.Println(i)
		select {
		case msg1 := <- c1:
			fmt.Println("received: ", msg1)
		case msg2 := <-c2:
			fmt.Println("recived2: ", msg2)
		}
	}

	fmt.Println(time.Since(now))


	//select {
	//case msg1 := <- c1:
	//	fmt.Println("received: ", msg1)
	//case msg2 := <-c2:
	//	fmt.Println("recived2: ", msg2)
	//}
}
