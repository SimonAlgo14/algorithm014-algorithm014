package main

import "fmt"

func combine(n int, k int) [][]int {
	var rs [][]int

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}

	cmb3(a, k, &rs, 0, 0, []int{})
	return rs
}

func cmb3(a []int, k int, rs *[][]int, left int, cnt int, one []int) {
	if cnt == k {
		*rs = append(*rs, append([]int{}, one...))
		return
	}

	if left >= len(a) {
		return
	}

	for i := left; i < len(a); i++ {
		one = append(one, a[i])

		cmb3(a, k, rs, i+1, cnt+1, one)

		one = one[:len(one)-1]
	}
}

func main() {
	a := combine(4, 2)
	fmt.Println(a)
}

/*
1 2 3 4
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
*/
