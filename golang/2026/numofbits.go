/* Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).

Example 1:

Input: 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
Example 2:

Input: 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit. */

package main

import "fmt"

func numberofBits(num uint32) int {
	ones := 0
	for num != 0 {
		if num%2 == 1 {
			ones++
		}
		num = num / 2
	}
	return ones
}

func main() {
	num := uint32(0b0000000000000000000000000000101)
	fmt.Println(numberofBits(num))
}
