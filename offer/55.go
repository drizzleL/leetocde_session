package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := maxDepth(root.Left), maxDepth(root.Right)
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

func isBalanced(root *TreeNode) bool {
	var helper func(root *TreeNode) (int, bool)
	helper = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		h1, ok1 := helper(root.Left)
		h2, ok2 := helper(root.Right)
		if !ok1 || !ok2 {
			return 0, false
		}
		if h2 > h1 {
			h1, h2 = h2, h1
		}
		if h1-h2 > 1 {
			return 0, false
		}
		// make h1 >= h2
		return 1 + h1, true
	}
	_, ok := helper(root)
	return ok
}
