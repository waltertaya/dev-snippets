#!/usr/bin/env python

from bst import TreeNode
from visualize import visualize_bst

if __name__ == "__main__":
    node = TreeNode(8)
    nodes = [3, 10, 1, 6, 14, 4, 7, 13, 12]
    for value in nodes:
        node.insert(value)
    # dot = visualize_bst(node)
    # dot.render("bst", format="png", view=True)
    print(node.inorder_traversal())
