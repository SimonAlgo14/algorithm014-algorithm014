package main

func water(a []int) int {
	var max int

	for i, j := 0, len(a); i < j; {
		var area int

		if a[i] < a[j] {
			area = a[i] * (j - i)
			i++
		} else {
			area = a[j] * (j - i)
			j--
		}

		if max < area {
			max = area
		}
	}

	return max
}
