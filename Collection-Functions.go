package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v ==t {
			return i
		}

	}
	return -1
}

func Include(vs []string, t string) bool  {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool  {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool  {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string  {
	vsf := make([]string, 0)

	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func Map(vs []string, f func(string) string) []string  {
	vsf := make([]string, len(vs))
	for i, v := range vs{
		vsf[i] = f(v)
	}
	return vsf
}

func main() {

	strs := []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "plum"))

	fmt.Println(Include(strs, "apple"))

	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "s")
	}))

	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "p")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
	fmt.Println(len(Map(strs, strings.ToUpper)))

}