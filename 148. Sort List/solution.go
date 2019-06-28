/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	l := 0
	for node := head; node != nil; node = node.Next {
		l++
	}
	return mergeList(head, l)
}
func mergeList(head *ListNode, l int) *ListNode {
	if l == 0 {
		return head
	} else if l == 1 {
		head.Next = nil
		return head
	}
	second := head
	for i := 0; i < l/2; i++ {
		second = second.Next
	}
	head = mergeList(head, l/2)
	second = mergeList(second, l-l/2)
	return merge(head, second)
}
func merge(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	node := dummy
	for ; l1 != nil || l2 != nil; node = node.Next {
		if l1 == nil {
			node.Next = l2
			break
		} else if l2 == nil {
			node.Next = l1
			break
		} else if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
	}
	return dummy.Next
}
