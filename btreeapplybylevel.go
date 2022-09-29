package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{})) (int, error) {
	if root == nil {
		return 1, nil
	}
	f(root.Data)
	right, left := root.Left, root.Right

	if left != nil {
		f(left)
		BTreeApplyByLevel(left, f)
	}
	if right != nil {
		f(right.Data)
		BTreeApplyByLevel(right, f)
	}
	return 0, nil
}
