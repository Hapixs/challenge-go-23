package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := BTreeLevelCount(root.Left), BTreeLevelCount(root.Right)
	return map[bool]int{true: left + 1, false: right + 1}[left > right]
}
