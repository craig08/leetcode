/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	for curr := head; curr != nil; curr = curr.Next {
		next := curr.Next
		for ; next != nil && next.Val == curr.Val; next = next.Next {
		}
		curr.Next = next
	}
	return head
}
