package main

import (
	"fmt"
	"time"
)

func main() {

	request := make(chan int, 5)

	for i:=0; i<5 ; i++  {
		request <- i
	}
	close(request)

	limiter := time.Tick(200*time.Millisecond)

	fmt.Println("start")

	for req := range request{
		<-limiter
		fmt.Println("request ",req, time.Now())
	}

	fmt.Println("end")


	burstyLimiter := make(chan time.Time, 3)

	for i:=0; i<3 ; i++  {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond*200){
			burstyLimiter <- t
		}
	}()


	burstyRequests := make(chan int, 5)

	for i:=0; i<5 ; i++  {
		burstyRequests <- 5
	}

	close(burstyRequests)

	for req := range burstyRequests {
		btime := <-burstyLimiter
		fmt.Println("request", req, time.Now(), btime)
	}

}