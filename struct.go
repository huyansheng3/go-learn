package main

import "fmt"

type Book struct {
	title string
	author string
}

func main()  {

	fmt.Println(Book{"Go 语言", "www.runoob.com",})


	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com"})


	fmt.Println(&Book{"Java编程实战", "小名"})

	b := Book{"Java编程实战", "小名"}

	fmt.Println(b.title)

	b2 := &b

	fmt.Println(b2.title)

	b2.author = `小龙`

	fmt.Println(b2)


}