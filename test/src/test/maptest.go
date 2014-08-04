package main

import (
	"fmt"
)

type Person struct {
	Id   string
	Name string
	Num  int32
}

type IError struct {
	Op   string
	Path string
	Err  error
}

func main() {
	var per map[string]Person
	per = make(map[string]Person, 5)
	per["123"] = Person{"456", "jack", 789}
	temp, ok := per["123"]
	if ok {
		tt := Person{temp.Id, "Tom", temp.Num}
		per["123"] = tt
	}
	fmt.Println(per["123"].Name)
	for i, pp := range per {
		fmt.Println(i, "tt", pp.Name)
	}

	var result []int
	result = append(result, 2, 5, 6)
	fmt.Println(result[0], len(result), cap(result))
}
