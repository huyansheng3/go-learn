package main

import "fmt"

func test()  {
	var var1 int = 6

	fmt.Printf("%T -> %v\n", var1, var1)

	var2 := float64(var1)

	var3 := int64(var1)


	fmt.Printf("%T -> %v\n", var2, var2)
	fmt.Printf("%T -> %v\n", var3, var3)


	var4 := new(int32)

	var5 :=(*int32)(var4)

	fmt.Printf("%T -> %v\n", var4, var4)

	fmt.Printf("%T -> %v\n", var5, var5)
}

func test2()  {
	var i interface{} = "kk"

	j, ok  := i.(int)

	if ok {
		fmt.Printf("%T -> %v\n", j, j)
	} else {
		fmt.Printf("%T -> %v\n", j, j)
	}


}

func main() {
	//test()
	test2()
}
