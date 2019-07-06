/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k-1 && node != nil; i++ {
		node = node.Next
	}
	if node == nil {
		return head
	}

	// reverse k nodes
	prev := reverseKGroup(node.Next, k)
	for n, i := head, 0; i < k; i++ {
		n.Next, prev, n = prev, n, n.Next
	}
	return prev
}
