package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42")
	} else {
		return arg +3, nil
	}
}

type argError struct {
	arg int
	prob string
}

func  (e argError)Error() string  {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error)  {
	if arg  == 42 {
		return -1, argError{arg, "cant work with it"}
	}
	return arg +3, nil
}

func main() {

	for _, i := range []int{7,42} {
		if r, e := f1(i);e!= nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 work", r)
		}
	}

	for _, i := range []int{7,42} {
		if r, e := f2(i);e!= nil {
			fmt.Println("&e", &e)
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 work", r)
		}
	}

	_, e := f2(42)

	fmt.Println("&e", &e)

	if ae, ok := e.(argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}


}