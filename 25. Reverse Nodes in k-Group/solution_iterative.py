# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        dummy = ListNode(0)
        dummy.next = head
        newHead = False
        preStart = dummy
        while preStart is not None:
            tail = preStart
            for i in range(k):
                tail = tail.next
                if tail is None:
                    return dummy.next

            node, nextPreStart = preStart.next, preStart.next
            preStart.next = tail.next
            for i in range(k):
                nextNode = node.next
                node.next = preStart.next
                preStart.next = node
                node = nextNode
            if not newHead:
                dummy.next = preStart.next
                newHead = True
            preStart = nextPreStart
        return dummy.next
