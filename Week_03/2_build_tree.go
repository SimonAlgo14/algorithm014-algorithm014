package main


func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return  nil
	}

	first := preorder[0]
	root  := TreeNode{ Left: nil, Right: nil, Val: first}

	var index int
	for index = 0; index < len(inorder); index++ {
		if inorder[index] == first {
			break
		}
	}

	root.Left  = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:],  inorder[index+1:])
	return &root
}

// preorder [root]   [left]   [right]
// inorder  [left]   root     [right]

//preorder = [3,     9,        20,15,7]
//inorder  = [9,     3,        15,20,7]