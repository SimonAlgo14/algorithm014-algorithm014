package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var cache []int

	return traversal(root, cache)
}


func traversal(root *TreeNode, cache []int) []int {
	if root == nil {
		return cache
	}

	cache = traversal(root.Left, cache)
	cache = append(cache, root.Val)

	return traversal(root.Right, cache)  //tail  recursive
}