package main

func levelOrder(root *TreeNode) [][]int {
	rs := make([][]int, 0)

	if root == nil {
		return rs
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var line []int
		var kids []*TreeNode

		for _, node := range queue {
			line = append(line, node.Val)

			if node.Left != nil {
				kids = append(kids, node.Left)
			}
			if node.Right != nil {
				kids = append(kids, node.Right)
			}
		}

		rs = append(rs, line)
		queue = kids
	}

	return rs
}

/*
    3
   / \
  9  20
    /  \
   15   7
*/
