package main

import (
	"fmt"
	"sync"
	"time"
)

func groupWorker(id int, wg *sync.WaitGroup)  {
	fmt.Printf("worker %d started!\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished!\n", id)

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	for i:=0;i< 5; i++  {
		wg.Add(1)
		go groupWorker(i, &wg)
	}

	wg.Wait()
}