package main

import (
	"fmt"
)

func sum(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}

func main() {
	ch := make(chan int, 0)
	arr := []int{1, 2, 3, 4, 5}
	go sum(arr, ch)
	arr1 := []int{5, 5, 5}
	go sum(arr1, ch)
	arr2 := []int{10, 10, 10}
	go sum(arr2, ch)
	x, y, z := <-ch, <-ch, <-ch
	fmt.Println(x + y + z) //o/p: 60
}
