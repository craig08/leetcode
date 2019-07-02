/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        ListNode* dummy = new ListNode(0);
        dummy->next = head;
        ListNode* node = dummy;
        // run example: 1->2->3->4->5, 2, 4
        // dummy->"1" (node) ->2->3->4->5
        for (; m > 1; --m, --n) {
            node = node->next;
        }
        // dummy->"1" (node) ->2->3->4->"5" (tail)
        ListNode* tail = node;
        for (; n >= 0; --n) {
            tail = tail->next;
        }
        node->next = reverse(node->next, tail);
        head = dummy->next;
        delete dummy;
        return head;
    }
    // reverse sub-list between [head, tail)
    ListNode* reverse(ListNode* head, ListNode* tail) {
        ListNode* prev = tail;
        // check tail instead of nullptr
        for (ListNode* node = head; node != tail; ) {
            ListNode* post = node->next;
            node->next = prev;
            prev = node;
            node = post;
        }
        return prev;
    }
};
