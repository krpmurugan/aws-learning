package main

import "fmt"

func main() {
	arr := []int{1, 4, 3, 7, 8, 9, 0, 6}
	tri, count, map1 := Triplet(arr, 7)
	fmt.Println(tri)
	fmt.Println(count)
	fmt.Println(map1)
}

func Triplet(arr []int, key int) (int, int, map[int][]int) {
	n := len(arr)
	if n < 3 {
		return -1, 0, nil // Array should have at least 3 elements
	}
	maxSum := 0
	count := 0
	map1 := make(map[int][]int)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := arr[i] + arr[j] + arr[k]
				if sum == key {
					maxSum = sum
					fmt.Printf("%d %d %d", arr[i], arr[j], arr[k])
					map1[i] = []int{arr[i], arr[j], arr[k]}
					count++
				}
			}
		}
	}
	return maxSum, count, map1
}
