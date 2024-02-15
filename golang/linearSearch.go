package main

import "fmt"

func LinearSearch(arr []int, key int) bool {
	for _, val := range arr {
		if val == key {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(LinearSearch(arr, 90))
}
