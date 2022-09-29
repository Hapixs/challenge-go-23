package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	right, left := root.Left, root.Right
	if left != nil {
		f(left.Data)
		BTreeApplyByLevel(left, f)
	}
	if right != nil {
		f(right.Data)
		BTreeApplyByLevel(right, f)
	}
}
