package main

func preorderTraversal(root *TreeNode) []int {
	var cache []int
	return preorder(root, cache)
}

func preorder(root *TreeNode, cache []int) []int{
	if root == nil {
		return cache
	}

	cache = append(cache, root.Val)
	cache = traversal(root.Left, cache)

	return traversal(root.Right, cache) //tail  recursive
}


