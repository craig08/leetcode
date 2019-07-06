/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy, newHead := &ListNode{Next: head}, false
	for preStart, tail := dummy, dummy; tail != nil; {
		for i := 0; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		prev := tail.Next
		for i, node := 0, preStart.Next; i < k; i++ {
			node.Next, prev, node = prev, node, node.Next
		}
		preStart.Next, preStart, tail = prev, preStart.Next, preStart.Next
		if !newHead {
			dummy.Next, newHead = prev, true
		}
	}
	return dummy.Next
}
