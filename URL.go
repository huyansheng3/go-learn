package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"


	 u, err := url.Parse(s)

	if err != nil{
		panic(err)
	}


	fmt.Println( u.Host )
	fmt.Println( u.Path )
	 fmt.Println(u.ForceQuery)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	fmt.Println(u.User)
	fmt.Println(u.Opaque)
	fmt.Println(u.Scheme)


 query, err :=url.ParseQuery(u.RawQuery)

	if err!=nil {
		panic(err)
	}

fmt.Println(query)

}