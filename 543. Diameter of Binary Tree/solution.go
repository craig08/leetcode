/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	diameter, _ := processNode(root)
	return diameter
}

func processNode(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	diaLeft, depLeft := processNode(root.Left)
	diaRight, depRight := processNode(root.Right)
	return max(diaLeft, diaRight, depLeft+depRight), max(depLeft, depRight) + 1
}

func max(nums ...int) int {
	ans := 0
	for _, num := range nums {
		if num > ans {
			ans = num
		}
	}
	return ans
}
