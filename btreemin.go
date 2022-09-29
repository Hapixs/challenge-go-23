package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data > root.Left.Data {
		return BTreeMin(root.Left)
	}
	if root.Data > root.Right.Data {
		return BTreeMin(root.Right)
	}
	return root
}
