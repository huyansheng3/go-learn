package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")

	fmt.Println(os.Getenv("FOO"))

	fmt.Println(os.Getenv("BAR"))

	for _,e := range os.Environ(){
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}