package main

import "fmt"

func main() {

	//for i:=0; i<15 ; i++  {
	//	fmt.Println(i)
	//}


	//j:= 0
	//
	//label:
	//if(j< 15) {
	//	fmt.Println(j)
	//	j+=1
	//	goto label
	//}


	//for i := 0; i < 5; i++ {
	//	var v int
	//	fmt.Printf("%d ", v)
	//	v = 5
	//}


	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}

	//i := 0
	//var j int
	//defer fmt.Println(j)
	//i+=1
	//j =100
	//return


	//test:= "胡衍生"
	//fmt.Println(len(test))

	//v := new(int)
	//
	//fmt.Println(v, *v)


	printNum(10)


}

func printNum(i int)  {
	if(i>0) {
		fmt.Println(i)
		printNum(i-1)
	}
}
