package main

import "fmt"

func findArrSum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func main() {
	fmt.Println(findArrSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
