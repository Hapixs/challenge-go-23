package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Parent == nil {
		*node = *root
	} else {
		*node = *node.Parent
	}
	return root
}
