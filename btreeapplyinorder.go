package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	f(root.Data)
}
