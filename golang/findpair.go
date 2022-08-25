package main

import "fmt"

func findPair(arr []int, sum int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Println("Find Pair", arr[i], arr[j])
			}
		}
	}
	return
}

func main() {
	arr1 := []int{10, 20, 30, 40, 50}
	findPair(arr1, 60)
}
