package main

import "fmt"

func BinarySearch(arr []int, needle int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		meadian := (low + high) / 2

		if arr[meadian] < needle {
			low = meadian + 1
		} else {
			high = meadian - 1
		}
	}

	if low == len(arr) || arr[low] != needle {
		return false
	}
	return true
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(BinarySearch(arr, 40))
}
