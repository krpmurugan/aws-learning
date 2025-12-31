/* 
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Input: [3,0,1]
Output: 2

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
*/

package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0
	for _, n := range nums {
		actualSum += n
	}

	// the difference between expected and actual is the missing number
	return expectedSum - actualSum
}

func main() {
  nums := []int{9,6,4,2,3,5,7,0,1}
  fmt.Println(missingNumber(nums))
}
