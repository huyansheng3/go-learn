package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		i := 0
		for  {
			fmt.Println("for ====  ", i)
			if i >3 {
				fmt.Println("if i: ", i)
				return
			} else {
				i+=1
			}
		}
	}()

	time.Sleep(time.Second)

	fmt.Println("test")

//	未经过channel receivce 阻塞主 goroutine, 主goroutine 执行结束后会结束所以的goroutine，所以goroutine执行结果每次都不一致

}