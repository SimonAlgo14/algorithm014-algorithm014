package main

import "fmt"

func main() {
	a := []int{ 5, 16, 27, 58, 0, 1, 3}
	fmt.Println( min(a), a[min(a)])

	b := []int{ 15, 0, 1, 3, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println( min(b), b[min(b)])
}

func min(a []int) int {
	lo, hi := 0, len(a)-1

	for lo < hi {
		if a[lo] < a[hi] {
			return lo
		}

		mi := lo + (hi-lo) / 2
		if a[mi] >= a[lo] {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return lo
}
