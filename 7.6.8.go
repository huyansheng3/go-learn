package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func FindFileDigits(filename string)[]byte  {
	var digitRegexp = regexp.MustCompile("[0-9]+")

	b ,_ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)

	c:= make([]byte, len(b))
	copy(c, b)
	return c
}
func main() {
	fmt.Println(string(FindFileDigits("./dup2-test")))

}
