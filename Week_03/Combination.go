package main

import "fmt"

func combination(a []rune, n int) {
	rs := make([]rune, n)
	cmb(a, rs, 0, 0, n)
}

func cmb(a []rune, rs []rune, s int, c int, n int) {
	if c == n {
		fmt.Println(string(rs))
		return
	}

	if s >= len(a) {
		return
	}

	for i := s; i < len(a); i++ {
		rs[c] = a[i]
		cmb(a, rs, i+1, c+1, n)
	}
}

func main() {
	combination([]rune("12345"), 5)
}
