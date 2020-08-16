package main

import "fmt"

func main() {
	arr := []int{0, 0, 0, 1, 0, 1, 0, 3, 12}
	fmt.Println(arr)

	moveZero(arr)
	fmt.Println(arr)
}

// [0, 1, 2]

func moveZero(a []int) {
	j := 0

	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			continue // case 1
		}

		a[j] = a[i] // case 2

		if i != j {
			a[i] = 0 // case 3
		}

		j++
	}
}