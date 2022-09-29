package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
}
