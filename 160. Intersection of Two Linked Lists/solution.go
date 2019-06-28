/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := findLen(headA), findLen(headB)
	if lenA > lenB {
		for ; lenA > lenB; lenA-- {
			headA = headA.Next
		}
	} else {
		for ; lenB > lenA; lenB-- {
			headB = headB.Next
		}
	}
	for headA != headB {
		headA, headB = headA.Next, headB.Next
	}
	return headA
}
func findLen(head *ListNode) int {
	l := 0
	for ; head != nil; head = head.Next {
		l++
	}
	return l
}
