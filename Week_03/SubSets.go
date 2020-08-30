package main

import (
	"fmt"
)

func subset(a []int) [][]int {
	var rs [][]int
	_subset(a, &rs, []int{}, 0, len(a)-1)
	return rs
}

func _subset(a []int, rs *[][]int, cache []int, l int, r int)  {
	if l > r {
		*rs = append(*rs, cache)
		return
	}

	_subset(a, rs, cache, l+1, r) // ignored

	cache = append(cache, a[l])
	_subset(a, rs, cache, l+1, r) // selected

	// cache = cache[:len(cache)-1]
}

func main() {
	fmt.Println(subset([]int{1, 2, 3}))
}
