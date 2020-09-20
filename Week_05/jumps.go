package main

import "fmt"

/*
   0 1 2 3 4
: [2,3,1,1,4]
   2
*/
func jump(nums []int) int {
	return jump_(nums, 0)
}

func jump_(nums []int, steps int) int {
	if len(nums) <= 0 {
		return steps
	}

	index := -1
	for i := len(nums) -1 ; i >= 0; i-- {
		if i + nums[i] >= len(nums) -1 {
			index = i
		}
	}

	fmt.Printf("steps: %d index:%d, %d, %v\n", steps, index, nums[index], nums[:index+1])

	if index == 0 {
		return steps + 1
	} else {
		return jump_(nums[:index+1], steps+1)
	}
}

func main() {
	a := []int{3, 1, 4, 2, 1, 1, 4}
	//a := []int{1}
	fmt.Println(jump(a))
}
