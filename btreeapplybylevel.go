package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)
	for _, q := range GetFromLevel(root) {
		f(q.Data)
	}
}

func GetFromLevel(root *TreeNode) []*TreeNode {
	queue := []*TreeNode{}
	if root == nil {
		return queue
	}
	if root.Left != nil {
		queue = append(queue, root.Left)
		if root.Left.Left != nil {
			queue = append(queue, GetFromLevel(root.Left)...)
		}
	}
	if root.Right != nil {
		queue = append(queue, root.Right)
	}
	return append(queue, GetFromLevel(root.Right)...)
}
