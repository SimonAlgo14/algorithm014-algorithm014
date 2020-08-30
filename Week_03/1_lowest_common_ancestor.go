package main


/*
                           3
                  5              1
              6       2       0        8
                  7       4
 */
func lowestCommonAncestor(root, p, q *TreeNode)  *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right,p, q)

	if l == nil && r == nil {
		return nil
	}

	if l != nil && r != nil {
		return root
	}

	if l != nil  {
		return l
	}

	return r
}
