/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	return construct(preorder, inorder, 0, len(preorder), 0, len(inorder))
}
func construct(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preEnd <= preStart || inEnd <= inStart {
		return nil
	}
	val, idx := preorder[preStart], inStart
	root := &TreeNode{Val: val}
	for ; idx < inEnd; idx++ {
		if inorder[idx] == val {
			break
		}
	}
	root.Left = construct(preorder, inorder, preStart+1, preStart+1+idx-inStart, inStart, idx)
	root.Right = construct(preorder, inorder, preStart+1+idx-inStart, preEnd, idx+1, inEnd)
	return root
}
