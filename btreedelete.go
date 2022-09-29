package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	*node = *node.Parent
	return root
}
