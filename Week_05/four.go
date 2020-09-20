package main

import (
	"fmt"
	"sort"
)

/*
nums = [1, 0, -1, 0, -2, 2]，target = 0。

-2 -1 0 0 1 2

[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */
func fourSum(nums []int, target int) [][]int {
	var r [][]int
	if len(nums) < 4 {
		return r
	}

	cache := make(map[string]bool)
	for i := 0; i < len(nums) - 1; i++ {
		arr := append([]int{}, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		ThreeSum(arr, target - nums[i], nums[i], &r, cache)
	}

	return r
}


func ThreeSum(nums []int, target int, first int, result *[][]int, cache map[string]bool) {
	if len(nums) < 3 {
		return
	}
	sort.Ints(nums)


	var low  int
	var high int
	var sum  int
	for i := 0; i < len(nums) - 1; i++ {
		low = i + 1
		high = len(nums) - 1
		sum = target - nums[i]
		for low < high {
			if nums[low] + nums[high]  ==  sum {
				one := []int{first, nums[i], nums[low], nums[high]}
				sort.Ints(one)

				key := fmt.Sprintf("%d-%d-%d-%d", one[0], one[1], one[2], one[3])
				if _, ok := cache[key]; !ok {
					cache[key] = true
					*result = append(*result, one)
				}

				low++
				high--

				//fmt.Println(result)
			} else if nums[low] + nums[high]  <  sum {
				low++
			} else {
				high--
			}
		}
	}
}

func main() {
	nums := [] int {1, 0, -1, 0, -2, 2}
	fmt.Println( fourSum(nums, 0))
}


