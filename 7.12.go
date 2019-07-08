package main

import "fmt"

func slice(str string, i int) []string  {
	res:=make([]string, 0)

	res = append(res, str[:i])
	res = append(res, str[i:])

	return res
}

func main() {
	//res := slice("abcde", 2)
	//fmt.Println(res)
	//
	//str :="abcd"
	//
	//fmt.Println(str[len(str)/2:] + str[:len(str)/2])


	//t:= "Google"
	//t:= "小胡是个帅哥"
	//
	//var x = []int32(t)
	//
	//for i:=0; i<len(x)/2 ;i++  {
	//	x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
	//	//fmt.Println(string(x[i]), string(x[len(x)-i-1]))
	//}
	//
	//fmt.Println(string(x))

	capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}


}
