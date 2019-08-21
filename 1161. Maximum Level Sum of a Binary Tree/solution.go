/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	q, ans, max, level := []*TreeNode{}, 0, 0, 1
	q = append(q, root)
	for len(q) > 0 {
		sum := 0
		for l := len(q); l > 0; l-- {
			n := q[0]
			q = q[1:]
			sum += n.Val
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		if sum > max {
			max, ans = sum, level
		}
		level++
	}
	return ans
}
