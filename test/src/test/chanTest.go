package main

import (
	"fmt"
	// "time"
)

func main() {
	ch := make(chan int)
	quit := make(chan int)
	go read(ch, quit)
	go write(ch)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch, "read")
	// }
	// time.Sleep(time.Second)
	// close(ch)
	<-quit
}

func write(b chan int) {
	for i := 0; i < 10; i++ {
		b <- i
		fmt.Println("write", i)
		// time.Sleep(time.Second)
	}
}

func read(b chan int, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-b, "read")
	}
	quit <- 0
}
