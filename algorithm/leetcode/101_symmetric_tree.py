class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        if root is None:
            return True
        return self.check(root.left, root.right)

    def check(self, a, b: TreeNode) -> bool:
        if (a is None or b is None) and a != b:
            return False
        return a.val == b.val and self.check(a.left, b.right) and self.check(a.right, b.left)
