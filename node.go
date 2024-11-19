package main

const MAX_TREE_HT = 100

type Node struct {
	data      rune
	frequency int
	left      *Node
	right     *Node
}

func newNode(data rune, freq int) *Node {
	node := &Node{
		left:      nil,
		right:     nil,
		data:      data,
		frequency: freq,
	}
	return node
}

func swapNodes(a **Node, b **Node) {
	t := *a
	*a = *b
	*b = t
}

func isLeaf(root *Node) bool {
	return root.left == nil && root.right == nil
}
