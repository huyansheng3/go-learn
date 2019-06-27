package main

import (
	"flag"
	"fmt"
	"time"
)

func fact(n int) int {
	if n == 0 || n == 1 {
		return n
	} else  {
		return fact(n-1) + fact(n-2)
	}
}


var number = flag.Int("n", 20, "fact n")

func main() {
	flag.Parse()

	now := time.Now()
	fmt.Println(fact(*number))
	spend := time.Since(now)
	fmt.Println(spend)

}
