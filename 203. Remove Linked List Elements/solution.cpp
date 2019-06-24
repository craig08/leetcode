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
    ListNode* removeElements(ListNode* head, int val) {
        ListNode* dummy = new ListNode(0);
        dummy->next = head;
        for (ListNode* node = dummy; node != nullptr; node = node->next) {
            while (node->next != nullptr && node->next->val == val) {
                node->next = node->next->next;
            }
        }
        head = dummy->next;
        delete dummy;
        return head;
    }
};
