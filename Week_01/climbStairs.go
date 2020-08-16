package main

func climbStairs(n int) int {
	x, y := 1, 2
	z    := 0

	if n <= 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	for i := 3; i <= n; i++ {
		z = x + y
		x = y
		y = z
	}

	return z
}

