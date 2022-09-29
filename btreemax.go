package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Right != nil {
		return BTreeMax(root.Right)
	}
	return root
}
