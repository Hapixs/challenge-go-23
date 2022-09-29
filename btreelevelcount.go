package piscine

func BTreeLevelCount(root *TreeNode) int {
	count := 0
	if root == nil {
		return 0
	}
	if root.Left != nil {
		count += 1 + BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		count += 1 + BTreeLevelCount(root.Right)
	}
	return count
}
