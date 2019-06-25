/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	for node := dummy; node.Next != nil && node.Next.Next != nil; node = node.Next.Next {
		node.Next = swap(node.Next)
	}
	return dummy.Next
}
func swap(head *ListNode) *ListNode {
	second := head.Next
	head.Next = second.Next
	second.Next = head
	return second
}
