/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}
	return construct(inorder, postorder, 0, len(inorder), 0, len(postorder))
}

func construct(inorder, postorder []int, inStart, inEnd, postStart, postEnd int) *TreeNode {
	if inEnd == inStart || postStart == postEnd {
		return nil
	}
	val, idx := postorder[postEnd-1], inStart
	root := &TreeNode{Val: val}
	for ; idx < inEnd && inorder[idx] != val; idx++ {
	}
	root.Left = construct(inorder, postorder, inStart, idx, postStart, postStart+idx-inStart)
	root.Right = construct(inorder, postorder, idx+1, inEnd, postStart+idx-inStart, postEnd-1)
	return root
}
