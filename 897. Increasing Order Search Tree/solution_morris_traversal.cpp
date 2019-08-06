/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* increasingBST(TreeNode* root) {
        TreeNode* dummy = new TreeNode(0);
        traversal(root, dummy);
        root = dummy->right;
        delete dummy;
        return root;
    }
    void traversal(TreeNode* root, TreeNode* arr) {
        for (TreeNode* node = root; node != nullptr;) {
            TreeNode* pred = findPred(node);
            if (pred == nullptr) {
                arr->right = node;
                arr = arr->right;
                node = node->right;
            } else if (pred->right == node) {
                pred->right = nullptr;
                node->left = nullptr;
                arr->right = node;
                arr = arr->right;
                node = node->right;
            } else {
                pred->right = node;
                node = node->left;
            }
        }
    }
    TreeNode* findPred(TreeNode* node) {
        if (node->left == nullptr) {
            return nullptr;
        }
        TreeNode* n = node->left;
        for (; n->right != nullptr && n->right != node; n = n->right) {}
        return n;
    }
};
