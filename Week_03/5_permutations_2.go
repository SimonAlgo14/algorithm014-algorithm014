package main

import (
	"fmt"
	"strconv"
	"strings"
)

func arrToStr(a []int ) string  {
	var ss []string
	for _, i := range a {
		ss = append(ss, strconv.Itoa(i))
	}

	return strings.Join(ss, "/")
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	rs := make([][]int, 0)
	hash := make(map[string]bool)
	permuteUnique_(nums, len(nums), &rs, hash,0)
	return rs
}

func permuteUnique_(a []int, k int, rs *[][]int, hash map[string]bool, left int) {
	if left == k {
		key := arrToStr(a)
		_, ok := hash[key]
		if ok {
			return
		}

		hash[key] = true
		*rs = append(*rs, append([]int{}, a...))

		return
	}

	for i := left; i < k; i++ {
		if i != left && a[i] == a[left] {
			continue
		}

		a[left], a[i] = a[i], a[left]

		permuteUnique_(a, k, rs, hash, left + 1) // left

		a[left], a[i] = a[i], a[left]
	}
}

func main()   {
	fmt.Println( permuteUnique([]int{2, 2, 1, 1}))

}

/*
input: [1,1,2]
output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]


[2,2,1,1]
[[1,1,2,2],[1,2,1,2],[1,2,2,1],[2,1,1,2],[2,1,2,1],[2,2,1,1]]

 */
