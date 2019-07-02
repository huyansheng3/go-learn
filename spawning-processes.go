package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOutput, err := dateCmd.Output()

	if err!=nil {
		panic(err)
	}

	fmt.Println(string(dateOutput))


	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}