package piscine

import "errors"

func BTreeApplyByLevel(root *TreeNode, f func(...interface{})) (int, error) {
	if root == nil {
		return 0, errors.New("root is nil")
	}
	f(root.Data)
	right, left := root.Left, root.Right

	if left != nil {
		f(left)
	}
	if right != nil {
		f(right.Data)
	}
	BTreeApplyByLevel(left, f)
	BTreeApplyByLevel(right, f)
	return 1, nil
}
