/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;

    Node() {}

    Node(int _val, Node* _prev, Node* _next, Node* _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};
*/
class Solution {
public:
    Node* flatten(Node* head) {
        Node* tail = nullptr;
        return flatten(head, tail);
    }
    Node* flatten(Node* head, Node* &tail) {
        for (Node* node = head; node;) {
            if (node->child) {
                Node* childTail = nullptr;
                Node* nextNode = node->next;
                node->next = flatten(node->child, childTail);
                node->next->prev = node;
                childTail->next = nextNode;
                if (nextNode != nullptr) {
                    nextNode->prev = childTail;
                }
                node->child = nullptr;
                node = nextNode;
                tail = childTail;
            } else {
                tail = node;
                node = node->next;
            }
        }
        return head;
    }
};
