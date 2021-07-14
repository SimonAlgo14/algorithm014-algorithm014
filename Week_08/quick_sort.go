package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenRandomArray(n int) []int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	a := make([]int, 0, n )

	for i := 0; i < n; i++ {
		a = append(a, r.Intn(100))
	}

	return a
}

func QuickSort(a []int, begin int, end int) {
	if end < begin {
		return
	}

	pivot := partition(a, begin, end)
	QuickSort(a, begin, pivot -1 )
	QuickSort(a, pivot  + 1 , end )
}

func partition(a []int, begin int, end int) int {
	pivot := end
	counter := begin

	for i := begin; i < end; i++ {
		if a[i] < a[pivot] {
			fmt.Println(i, counter, a)
			a[i], a[counter] = a[counter], a[i]
			counter++
		}
	}

	a[counter], a[pivot] = a[pivot], a[counter]
	fmt.Println(counter, pivot, a)

	return counter
}

func main() {
	a := GenRandomArray(20)
	fmt.Println(a)
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
