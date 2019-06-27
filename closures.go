package main

import "fmt"

func increase() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := increase()

	fmt.Println(nextInt())

	fmt.Println(nextInt())
	fmt.Println(nextInt())


	nextInts := increase()

	fmt.Println(nextInts())


}
