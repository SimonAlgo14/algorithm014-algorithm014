package main

import "fmt"

func main() {
	a := []int{0, 0, 0, 0, 0, 1, 1, 2, 3, 4, 4, 5}
	n := removeDuplicates(a)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])

	}
	fmt.Println()
}

func removeDuplicates(a []int) int {
	if a == nil || len(a) == 0 {
		return 0
	}

	j := 1
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			a[j] = a[i]
			j++
		}
	}

	return j
}
