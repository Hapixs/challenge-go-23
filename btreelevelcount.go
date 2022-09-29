package piscine

func BTreeLevelCount(root *TreeNode) int {
	count := 1
	if root == nil {
		return 0
	}
	if root.Left != nil {
		count += BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		count += BTreeLevelCount(root.Right)
	}
	return count
}
