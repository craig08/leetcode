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
    ListNode* rotateRight(ListNode* head, int k) {
        if (!head) return head;
        ListNode* fast = head;
        ListNode* slow = head;
        int len = 0;
        for (; fast; ++len) {
            fast = fast->next;
        }
        k %= len;
        fast = head;
        for (; k > 0; --k, fast = fast->next) {}
        for (; fast->next; fast = fast->next, slow = slow->next) {}
        fast->next = head;
        head = slow->next;
        slow->next = nullptr;
        return head;
    }
};
