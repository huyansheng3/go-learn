package main

import (
	"fmt"
	"time"
)

/**
干活的人，可以有好几个人一起干活
 */
func worker1(id int, jobs <- chan int, result chan<- int)  {
	for i := range jobs{
		fmt.Println("worker", id, "start job")
		time.Sleep(time.Second * 2)
		fmt.Println("worer", id, "finish job")
		result <- i
	}
}

func main() {

	jobs := make(chan int, 100)
	results:= make(chan int, 100)

	fmt.Println("prepare worker")
	for w:=0; w<3 ; w++  {
		go worker1(w, jobs, results)
	}

	fmt.Println("prepare jobs")
	for j:=0; j< 10 ; j++  {
		fmt.Println("send job: ", j)
		jobs <- j
	}

	fmt.Println("get results")
	for i:=0;i<10 ;i++  {
		//只有receive chan 才会阻塞，相当于阻塞了for循环，有点神奇
		//<-results
		fmt.Println("result: ", <-results)
	}
}

