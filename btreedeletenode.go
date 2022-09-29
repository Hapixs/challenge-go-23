package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Parent == nil {
		*node, *node.Right = *root.Left, *root.Right
	} else {
		*node = *node.Parent
	}
	return root
}
