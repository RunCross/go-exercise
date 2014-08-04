package main

import "fmt"
import "./oop"

func main() {
	t := &oop.Oop{"int", "654"}
	var temp *oop.Oop
	fmt.Println(t)
	fmt.Println(temp)
}
