package main

import "fmt"

func main() {

	arr := []int{90, 70, 30, 30, 10, 120, 40, 150, 40, 30}
	var n int
	for j := 0; j < len(arr); j++ {
		if n < arr[j] {
			n = arr[j]
		}
		//fmt.Println(arr)
	}
	fmt.Println(n)
}
