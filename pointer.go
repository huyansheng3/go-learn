package main

import (
	"fmt"
	"strings"
)

func zeroVal(val int)  {
	val = 0
}

func zeroPtr(val *int)  {
	*val = 0
}


type person struct {
	name string
	age int
}


type ByteSize float64
const (
	_ = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {



	test2 := 1

	zeroVal(test2)

	fmt.Println("test2", test2)


	test := 1

	zeroPtr(&test)

	fmt.Println("test", test)

	fmt.Println("&test", &test)

	p1:= person{"hys", 12}

	fmt.Printf("%v\n", &p1)

	//b := true

	// 类型不同无法进行比较，bool与int无法进行转型
	//if int(b) == test {
	//	fmt.Printf("dsds")
	//}

	//fmt.Println(KB,
	//MB,
	//GB,
	//TB,
	//PB,
	//EB,
	//ZB,
	//YB)

	fmt.Println(strings.IndexRune("chicken", rune('k')))




}
