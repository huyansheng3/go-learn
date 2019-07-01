package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (l byLength) Len() int  {
	return len(l)
}

func (l byLength)Less(i, j int) bool  {
	return len(l[i]) < len(l[j])
}

func  (l byLength)Swap(i, j int)  {
	l[i], l[j] = l[j], l[i]
}

func main() {
	fruits := []string{"apple", "banana", "lemon", "cherry", "z"}

	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
