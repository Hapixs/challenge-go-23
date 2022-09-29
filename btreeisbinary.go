package piscine

func BtreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		if root.Right.Data < root.Left.Data {
			return false
		}
		return BtreeIsBinary(root.Left) && BtreeIsBinary(root.Right)
	}
	return true
}
