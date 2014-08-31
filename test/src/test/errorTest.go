package main

import (
	"fmt"
)

type iError struct {
	Op   string
	Path string
	Err  error
}

func (e *iError) Error() string {
	return "my error"
}

func ErrorTest() error {
	return &iError{"op", "path", nil}
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	err := ErrorTest()
	if op, ok := err.(*iError); ok {
		fmt.Println(op.Op, op.Path, op.Err)
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	if ae, ok := e.(*iError); ok {
		fmt.Println(ae.Op)
		fmt.Println(ae.Path)
	} else {
		fmt.Println("NOT IS iError")
	}

	if _, e := f2(42); e != nil {
		fmt.Println(e.Error())
	}
}
