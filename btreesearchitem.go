package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		if root.Left.Data == elem {
			return root.Left
		}
		rt := BTreeSearchItem(root.Left, elem)
		if rt != nil {
			return rt
		}
	}
	if root.Right != nil {
		if root.Right.Data == elem {
			return root.Right
		}
		rt := BTreeSearchItem(root.Right, elem)
		if rt != nil {
			return rt
		}
	}
}
