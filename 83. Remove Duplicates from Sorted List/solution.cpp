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
    ListNode* deleteDuplicates(ListNode* head) {
        for (ListNode* node = head; node != nullptr; node = node->next) {
            while (node->next != nullptr && node->next->val == node->val) {
                node->next = node->next->next;
            }
        }
        return head;
    }
};
