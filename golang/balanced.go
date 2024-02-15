package main

import "fmt"

func Balanced(text string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(text); i++ {
		if text[i] == '(' || text[i] == '{' || text[i] == '[' {
			stack = append(stack, text[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if text[i] == ')' && top != '(' || text[i] == '}' && top != '{' || text[i] == ']' && top != '[' {
				return false
			}
		}
	}
	return true
}

func main() {
	text := "({[})"
	a := Balanced(text)
	fmt.Println(a)
}
