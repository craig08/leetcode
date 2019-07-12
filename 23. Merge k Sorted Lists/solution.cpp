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
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        auto comp = [](ListNode* a, ListNode* b) {
            return a->val > b->val;
        };
        priority_queue<ListNode*, vector<ListNode*>, decltype(comp)> pq(comp);
        for (ListNode* list: lists) {
            if (list == nullptr) {
                continue;
            }
            pq.push(list);
        }
        ListNode* dummy = new ListNode(0);
        ListNode* tail = dummy;
        while (!pq.empty()) {
            ListNode* node = pq.top();
            pq.pop();
            if (node->next != nullptr) {
                pq.push(node->next);
            }
            tail->next = node;
            tail = tail->next;
        }
        tail = dummy->next;
        delete dummy;
        return tail;
    }
};
