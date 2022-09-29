package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		return BTreeMax(root.Left)
	}
	return root
}
