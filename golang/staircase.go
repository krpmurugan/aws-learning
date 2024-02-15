package main

import "fmt"

func StairCase(input int) int {
	if input == 0 {
		return 0
	}
	if input == 1 {
		return 1
	}
	if input == 2 {
		return 2
	}

	res := StairCase(input-1) + StairCase(input-2)
	fmt.Println(res)
	return res

}

func StairCase1() {
	input := 5
	for i := 0; i < input; i++ {
		for j := 0; j < input-i-1; j++ {
			fmt.Print("")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	n := 5
	res := StairCase(n)
	fmt.Println(res)

	StairCase1()

}
