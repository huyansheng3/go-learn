package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	go func() {
		for  {
			fmt.Println("for ----")
			j, more := <- jobs

			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				//return
			}
		}
	}()

	for j:=1; j<=3 ; j++  {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)

	fmt.Println("sent all jobs")

	d1 := <- done

	fmt.Println("done: ", d1)

}