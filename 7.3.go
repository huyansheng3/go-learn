package main

import "fmt"

func main() {
	//s:=make([]byte, 5)
	//
	//s=s[2:4]
	//
	//fmt.Println(len(s), cap(s))

	//s1 := []byte{'p', 'o', 'e', 'm'}
	//
	//s2:= s1[2:]
	//fmt.Println(s2)
	//
	//s2[1] = byte(112)
	//
	//fmt.Println(s2)
	//
	//fmt.Println(s1)

	items := [...]int{10, 20, 30, 40, 50}

	//for _, item := range items {
	//	item *= 2
	//}

	for i:=0; i<len(items) ; i++  {
		items[i] *=2
	}

	fmt.Println(items)



}
