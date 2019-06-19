package main

import "fmt"

type Book struct {
	title string
	author string
	subject string
	book_id int
}

func main()  {

	fmt.Println(Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})


	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})


}