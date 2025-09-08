from graphviz import Digraph

def visualize_bst(node):
    dot = Digraph()
    
    def add_nodes_edges(node):
        if node:
            dot.node(str(id(node)), str(node.value))
            if node.left:
                dot.edge(str(id(node)), str(id(node.left)))
                add_nodes_edges(node.left)
            if node.right:
                dot.edge(str(id(node)), str(id(node.right)))
                add_nodes_edges(node.right)
    add_nodes_edges(node)
    return dot
