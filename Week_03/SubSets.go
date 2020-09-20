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
		*rs = append(*rs, append([]int{}, cache...))  // **** attention  must copy **
		return
	}

	_subset(a, rs, cache, l+1, r) // ignored
	_subset(a, rs, append(cache, a[l]), l+1, r) // selected

	// cache = cache[:len(cache)-1]
}

func main() {
	fmt.Println(subset([]int{1, 2, 3}))
}


func bfs(nums []int) [][]int {
	var r [][]int
	if len(nums) <= 0 {
		return r
	}

	r = append(r, []int{})
	for _, num := range nums {
		var t [][]int
		for _, one := range r {
			two := append([]int{}, one...)  // clone, required !!!
			t = append(t, append(two, num))
		}
		r = append(r, t...)
	}
	return r
}