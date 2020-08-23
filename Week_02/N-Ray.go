package main

type NNode struct {
	Val      int
	Children []*NNode
}

func npreorder(root *NNode) []int {
	var cache []int

	return visitor(root, cache)
}

func visitor(root *NNode, cache []int) []int {
	if root == nil {
		return cache
	}

	cache = append(cache, root.Val)

	for _, child := range root.Children {
		cache = visitor(child, cache)
	}

	return cache
}