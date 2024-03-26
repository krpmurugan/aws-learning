package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 20, 40, 30}
	fmt.Println(maxTriplet(arr))
}

func maxTriplet(arr []int) int {
	n := len(arr)
	if n < 3 {
		return -1
	}

	maxSum := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := arr[i] + arr[j] + arr[k]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}
