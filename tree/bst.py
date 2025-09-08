class TreeNode:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right

    def insert(self, value):
        if value < self.value: # goto the left
            if self.left:
                self.left.insert(value)
            else:
                self.left = TreeNode(value)
        elif value > self.value: # goto the right
            if self.right:
                self.right.insert(value)
            else:
                self.right = TreeNode(value)
        else: # value same as self.value or handle duplicates
            pass
    
    def inorder_traversal(self):
        result = []
        if self.left:
            result.extend(self.left.inorder_traversal())
        result.append(self.value)
        if self.right:
            result.extend(self.right.inorder_traversal())
        return result
