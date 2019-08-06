# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        while root:
            # if the node has left child, insert left sub-tree into its right child
            if root.left:
                pred = root.left
                while pred.right:
                    pred = pred.right
                pred.right, root.left, root.right = root.right, None, root.left
            root = root.right
