package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	solution := 0
	l := len(s)
	for i := 0; i < l; i++ {
		var substring int
		letters := make(map[string]bool, 0)

		letters[string(s[i])] = true
		substring += 1

		for j := i + 1; j < l; j++ {
			if _, ok := letters[string(s[j])]; ok {
				break
			} else {
				substring += 1
				letters[string(s[j])] = true
				fmt.Println(substring)
			}
		}

		if substring > solution {
			solution = substring
		}
	}
	return solution
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aabccbb"))
}
