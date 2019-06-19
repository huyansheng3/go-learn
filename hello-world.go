package main

import "fmt"

func main()  {
	//var a = "RUN"
	//fmt.Println(a)
	//
	//var b int
	//fmt.Println(b)
	//
	//var c bool
	//
	//fmt.Println(c)

	//const (
	//	a = iota   //0
	//	b          //1
	//	c          //2
	//	d = "ha"   //独立值，iota += 1
	//	e          //"ha"   iota += 1
	//	f = 100    //iota +=1
	//	g          //100  iota +=1
	//	h = iota   //7,恢复计数
	//	i          //8
	//)
	//fmt.Println(a,b,c,d,e,f,g,h,i)


	var a uint = 60      /* 60 = 0011 1100 */
	var b uint = 13      /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b       /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c )

	c = a | b       /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c )

	c = a ^ b       /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c )

	c = a << 2     /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c )

	c = a >> 2     /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c )

}