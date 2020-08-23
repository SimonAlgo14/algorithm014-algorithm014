package main

import "fmt"

func min2(x int, y int)  int {
	if x < y {
		return x
	}
	return y
}

func min(x int, y int, z int)  int {
	t := min2(x, y)
	return min2(t, z)
}

func nthUglyNumber(n int) []int {
	i2, i3, i5 := 0, 0, 0
	f2, f3, f5 := 2, 3, 5
	a := make([]int, n)
	a[0] = 1

	for j := 1; j < n; j++ {
		f235 := min(f2, f3, f5)
		a[j] = f235

		if f235 == f2 {
			i2 ++
			f2 = a[i2] * 2
		}

		if f235 == f3 {
			i3 ++
			f3 = a[i3] * 3
		}

		if f235 == f5 {
			i5 ++
			f5 = a[i5] * 5
		}
	}

	return a
}

func main() {
	fmt.Println( nthUglyNumber(34 ))
}

