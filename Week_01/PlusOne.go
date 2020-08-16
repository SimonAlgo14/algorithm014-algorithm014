package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", PlusOne([]int{9, 9}))
	fmt.Printf("%v\n", PlusOne([]int{8, 9, 9}))
	fmt.Printf("%v\n", PlusOne([]int{9, 9, 8, 9, 9}))
}

func PlusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int {}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
