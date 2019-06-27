package main

import "fmt"

func zeroVal(val int)  {
	val = 0
}

func zeroPtr(val *int)  {
	*val = 0
}

func main() {



	test2 := 1

	zeroVal(test2)

	fmt.Println("test2", test2)


	test := 1

	zeroPtr(&test)

	fmt.Println("test", test)

	fmt.Println("&test", &test)


}
