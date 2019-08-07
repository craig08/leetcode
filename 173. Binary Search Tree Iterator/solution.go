/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	Node *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	if root == nil {
		return BSTIterator{}
	}
	return BSTIterator{Node: traverse(root)}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	num := this.Node.Val
	this.Node = traverse(this.Node.Right)
	return num
}

func traverse(node *TreeNode) *TreeNode {
	for node != nil {
		if node.Left != nil {
			pred := node.Left
			for ; pred.Right != nil && pred.Right != node; pred = pred.Right {
			}
			if pred.Right == nil {
				pred.Right, node = node, node.Left
			} else {
				pred.Right = nil
				return node
			}
		} else {
			return node
		}
	}
	return node
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.Node != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
