package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		if root.Right.Data < root.Left.Data {
			return false
		}
		return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
	}
	return true
}
