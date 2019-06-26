/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	fast, slow, tail := head, head, head
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		tail = slow
		slow = slow.Next
	}
	tail.Next = nil
	slow = reverse(slow)
	for fast = head; fast != nil && slow != nil; fast = fast.Next.Next {
		nextSlow := slow.Next
		slow.Next = fast.Next
		fast.Next = slow
		slow = nextSlow
		if fast.Next == nil {
			break
		}
	}
}
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for node := head; node != nil; {
		post := node.Next
		node.Next = prev
		prev = node
		node = post
	}
	return prev
}
