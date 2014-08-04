package main

import "fmt"
import "./oop"

func main() {
	t := &oop.Oop{"int", "654"}
	fmt.Println(t)
	fmt.Println(t.Add())
}
