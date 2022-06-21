1. Write a program that prints the numbers from 1 to 100. But for multiples of three print, “Fizz” instead of the number, and for multiples of five, print “Buzz”. For numbers which are multiples of both three and five, print “FizzBuzz”.

```
package main

import "fmt"

func main() {
	const (
		FIZZ = 3 
		BUZZ = 5
	)
	var p bool                 
	for i := 1; i < 100; i++ { 
		p = false
		if i%FIZZ == 0 { 
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 { 
			fmt.Printf("Buzz")
			p = true
		}
		if !p { 
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
```

2. Write a function that performs a bubble sort on a slice of ints?
```
package main

import (
	"fmt"
)

func main() {
	n := []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("unsorted %v\n", n)
	bubblesort(n)
	fmt.Printf("sorted %v\n", n)
}

func bubblesort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}
```

3. Write a function that finds the maximum value in an int slice ([]int).

```
func max(l []int) (max int) {  
    max = l[0]
    for _, v := range l {   
        if v > max {    
            max = v
        }
    }
    return
}

```

4. Paranthesis check Program

```
package main

import "fmt"

func main() {
	s := "({}[])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	// set up stack and map
	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// loop over string
	for _, r := range s {

		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		st = st[:l]
	}

	return len(st) == 0
}

```
