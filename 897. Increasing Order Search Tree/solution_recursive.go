/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = increasingBST(root.Right)
	left := increasingBST(root.Left)
	if left == nil {
		return root
	}
	node := left
	for ; node.Right != nil; node = node.Right {
	}
	node.Right, root.Left = root, nil
	return left
}
