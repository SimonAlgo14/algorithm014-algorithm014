package main

import "fmt"

func permutation(a []rune) {
	p(a, 0, len(a)-1)
}

//permute the value from left to right
func p(a []rune, l int, r int) {
	if l == r {
		fmt.Println(string(a))
		return
	}

	for i := l; i <= r; i++ {
		a[l], a[i] = a[i], a[l]
		p(a, l+1, r)
		a[l], a[i] = a[i], a[l]
	}
}

func main() {
	permutation([]rune("123"))
}
