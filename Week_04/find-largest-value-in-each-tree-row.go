package main

import "math"

func largestValues(root *TreeNode) []int {
	q := []*TreeNode{ root }
	r := make([]int, 0)

	if root == nil {
		return r
	}

	for len(q) > 0 {
		c := make([]*TreeNode, 0)

		max := math.MinInt32
		for _, one := range q {
			if one.Val > max {
				max = one.Val
			}
			if one.Left != nil {
				c = append(c, one.Left)
			}
			if one.Right != nil {
				c = append(c, one.Right)
			}
		}

		r = append(r, max)
		q = c
	}

	return  r
}

