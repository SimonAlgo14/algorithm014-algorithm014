package main

import "fmt"

func pow(x float64, n int) float64 {
	if n < 0 {
		n = -1 * n
		x = 1 / x
	}

	return _pow(x, n)
}

func _pow(x float64, n int) float64 {
	if n == 0 {
		return  1.0
	}

	if n == 1 {
		return x
	}

	var adj float64
	if n % 2 == 0 {
		adj = 1
	} else {
		adj = x
	}

	half := _pow(x, n / 2)
	return adj * half * half
}

func main() {
	fmt.Println(pow(2.0, 10))
	fmt.Println(pow(2.0, -3))
}
