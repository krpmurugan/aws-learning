package main

// Count the number of Digits in a number

import "fmt"

func main() {
	var n int
	count := 0
	n = 123456
	for n > 0 {
		n = n / 10
		count++
	}
	fmt.Println(count)
}
