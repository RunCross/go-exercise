package main

import (
	"fmt"
)

func main() {
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// var arr []int
	// arr = make([]int, 10)
	// for i, _ := range arr {
	// 	arr[i] = i
	// }
	// fmt.Println(arr)
	// chang(arr)
	// fmt.Println(arr)
	arr := []int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println(arr)
}

func modify(arr []int) {
	arr[0] = 999999
}

func chang(arr []int) {
	arr[5] = 9
}
