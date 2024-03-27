package main

/*Problem Statement
Given a string, find the length of the longest substring which has no repeating characters.

Example 1:

Input: String="aabccbb"
Output: 3
Explanation: The longest substring without any repeating characters is "abc".

Example 2:

Input: String="abbbb"
Output: 2
Explanation: The longest substring without any repeating characters is "ab".

Example 3:

Input: String="abccde"
Output: 3
Explanation: Longest substrings without any repeating characters are "abc" & "cde".*/

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

func longestSubstring(str string) int {
	m := make(map[byte]byte)
	count := 0
	max := 0
	s := []byte{}
	res := []byte{}
	for i := 0; i < len(str); i++ {
		if _, ok := m[str[i]]; ok {
			if count > max {
				max = count
				res = s
			}
			count = 1
			s = []byte{}
		} else {
			m[str[i]] = str[i]
			s = append(s, str[i])
			count++
		}
	}
	fmt.Println(string(res))
	return max
}


func main() {
	fmt.Println(lengthOfLongestSubstring("aabccbb"))
}
