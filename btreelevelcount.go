package piscine

func BTreeLevelCount(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}
	if root.Left != nil {
		count += BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		count += BTreeLevelCount(root.Right)
	}
	return count
}
