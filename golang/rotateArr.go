package main

import "fmt"

func RotateArray(arr []int, rotate int, size int) []int {
	var newArray []int
	for i := 0; i < rotate; i++ {
		newArray = arr[1:size]
		newArray = append(newArray, arr[0])
		arr = newArray
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(RotateArray(arr, 2, len(arr)))
}
