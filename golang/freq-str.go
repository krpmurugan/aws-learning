package main

import "fmt"

func main() {
	arr := "geeksforgeeks"
	freq := make(map[string]int)
	for _, char := range arr {
		freq[string(char)] = freq[string(char)] + 1
	}
	fmt.Println(freq)
}
