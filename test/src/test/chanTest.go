package main

import (
	"fmt"
)

func main() {
	ch := make(chan []byte)
	go test(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch, "read")
	}
}

func test(b chan []byte) {
	for i := 0; i < 10; i++ {
		c := make([]byte, 2)
		c[0] = 2
		c[1] = 3
		b <- c
		fmt.Println("write")
	}
}
