package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := [3]string{"a", "e", "D"}

	strSlice := strs[:]

	sort.Strings(strSlice)

	fmt.Println("sorted strings", strs, strSlice)

	fmt.Println(&strs[0], &strSlice[0])

	//fmt.Println("===", strSlice == strs)

	ints := []int{11, 3, 43, 3}
	sort.Ints(ints)
	fmt.Println("sorted ints", ints)

	sorted := sort.IntsAreSorted(ints)

	fmt.Println("sorted", sorted)

}