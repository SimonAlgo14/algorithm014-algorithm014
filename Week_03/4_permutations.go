package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	rs := make([][]int, 0)
	permute_(nums, len(nums), &rs, 0)
	return rs
}

func permute_(a []int, k int, rs *[][]int, left int) {
	if left == k {
		*rs = append(*rs, append([]int{}, a...))
		return
	}

	for i := left; i < k; i++ {
		a[left], a[i] = a[i], a[left]

		permute_(a, k, rs, left + 1) // left

		a[left], a[i] = a[i], a[left]
	}
}

func main()   {
	fmt.Println( permute([]int{1, 2, 3}))

}

