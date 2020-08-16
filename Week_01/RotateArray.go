package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fastRotate(a, 3)

	a = []int{1, 2, 3, 4, 5, 6, 7}
	slowRotate(a, 3)
}

func fastRotate(a []int, k int)  {
	nk := k % len(a)

	fmt.Printf("%v\n", a)
	reverse(a, 0, len(a) - 1)

	fmt.Printf("%v\n", a)
	reverse(a, 0, nk- 1)

	fmt.Printf("%v\n", a)
	reverse(a, nk, len(a) - 1)

	fmt.Printf("\n\nFast: %v\n", a)
}

func reverse(a []int, s int, e int)  {
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}
}

func slowRotate(a []int, k int)  {
	length := len(a)
	nk := k % len(a)
	for i := 0; i < nk; i++ {
		last := a[length-1]
		for j := length - 1; j > 0; j-- {
			a[j] = a[j-1]
		}
		a[0] = last
	}

	fmt.Printf("\n\nSlow: %v\n", a)
}