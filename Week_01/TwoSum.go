package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", TwoSum([]int{2, 7, 11, 15}, 9))
}

func TwoSum(nums []int, target int) []int {
	h := make(map[int]int, len(nums))  // reduce memory usage

	for i, n := range nums {
		if _, ok := h[n]; ok {
			return []int{h[n], i}
		}
		h[target-n] = i
	}

	return nil
}
