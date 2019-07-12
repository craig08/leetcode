/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, ok := balanced(root)
	return ok
}

func balanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	var left, right int
	var ok bool
	left, ok = balanced(root.Left)
	if !ok {
		return 0, false
	}
	right, ok = balanced(root.Right)
	if !ok {
		return 0, false
	}
	if left-right < -1 || left-right > 1 {
		return 0, false
	}
	if left > right {
		return left + 1, true
	}
	return right + 1, true
}
