/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{0, head}
	preNode := dummy
	for i := 1; i < m; i++ {
		preNode = preNode.Next
	}
	node, tail := preNode.Next, preNode.Next
	preNode.Next = nil
	for i := m; i <= n; i++ {
		nextNode := node.Next
		node.Next = preNode.Next
		preNode.Next = node
		node = nextNode
	}
	tail.Next = node
	return dummy.Next
}
