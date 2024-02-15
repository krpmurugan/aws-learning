package main

import "fmt"

func chunkSize(slice []int, chunckSize int) [][]int {
	var chunks [][]int
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < chunckSize {
			chunckSize = len(slice)
		}
		chunks = append(chunks, slice[0:chunckSize])
		slice = slice[chunckSize:]
	}
	return chunks
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	chunks := chunkSize(arr, 3)
	fmt.Println(chunks)
}
