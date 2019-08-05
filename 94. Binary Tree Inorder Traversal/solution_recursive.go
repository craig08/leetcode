/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(&result, root)
	return result
}

func inorder(result *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	inorder(result, node.Left)
	*result = append(*result, node.Val)
	inorder(result, node.Right)
}
