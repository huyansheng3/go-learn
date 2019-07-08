package main

import (
	"fmt"
)

func factN(n int) (int ){
	if n ==1 {
		 return 1
	}
	return n * factN( n-1 )
}

func main() {

	//start := time.Now()
	//
	//fmt.Println(factN(16))
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Printf("longCalculation took this amount of time: %s\n", delta)


	//a := [...]string{"a", "b", "c", "d"}
	//for i := range a {
	//	fmt.Println("Array item", i, "is", a[i])
	//}

	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	//// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice
	//
	//for i:=0; i < len(arrKeyValue); i++ {
	//	fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	//}

	b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4])
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])
}