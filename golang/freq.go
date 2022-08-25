package main

import "fmt"

func main() {
	arr := []int{90, 70, 30, 30, 10, 80, 40, 50, 40, 30}
	freq := make(map[int]int)
	for _, fre := range arr {
		freq[fre] = freq[fre] + 1
	}
	fmt.Println(freq)
}
