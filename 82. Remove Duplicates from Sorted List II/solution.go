/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	for node := dummy; node != nil && node.Next != nil; {
		val, duplicate, follow := node.Next.Val, false, node.Next.Next
		for ; follow != nil && follow.Val == val; follow = follow.Next {
			duplicate = true
		}
		if duplicate {
			node.Next = follow
		} else {
			node = node.Next
		}
	}
	return dummy.Next
}
