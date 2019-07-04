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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        stack<int> s1, s2;
        for (; l1 != nullptr; l1 = l1->next) {
            s1.push(l1->val);
        }
        for (; l2 != nullptr; l2 = l2->next) {
            s2.push(l2->val);
        }
        int carry = 0;
        ListNode* node = nullptr;
        for (; !s1.empty() || !s2.empty() || carry > 0; ) {
            if (!s1.empty()) {
                carry += s1.top();
                s1.pop();
            }
            if (!s2.empty()) {
                carry += s2.top();
                s2.pop();
            }
            ListNode* prev = new ListNode(carry % 10);
            carry /= 10;
            prev->next = node;
            node = prev;
        }
        return node;
    }
};
