package main

func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val && isSub(a, b) {
		return true
	}
	return isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}

func isSub(a *TreeNode, b *TreeNode) bool {
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	if b.Left != nil {
		if !isSub(a.Left, b.Left) {
			return false
		}
	}
	if b.Right != nil {
		if !isSub(a.Right, b.Right) {
			return false
		}
	}
	return true
}
