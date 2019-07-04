/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1, len2 := getLen(l1), getLen(l2)
	if len2 > len1 {
		l1, l2, len1, len2 = l2, l1, len2, len1
	}
	var node *ListNode
	for l1 != nil && l2 != nil {
		var prev *ListNode
		if len1 > len2 {
			prev = &ListNode{l1.Val, node}
		} else {
			prev = &ListNode{l1.Val + l2.Val, node}
			l2, len2 = l2.Next, len2-1
		}
		node = prev
		l1, len1 = l1.Next, len1-1
	}
	return reverse(node)
}
func getLen(l *ListNode) (n int) {
	for ; l != nil; n, l = n+1, l.Next {
	}
	return
}
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	carry := 0
	for node := head; node != nil; {
		sum := carry + node.Val
		node.Val, carry = sum%10, sum/10
		post := node.Next
		node.Next = prev
		prev, node = node, post
	}
	if carry > 0 {
		return &ListNode{carry, prev}
	}
	return prev
}
