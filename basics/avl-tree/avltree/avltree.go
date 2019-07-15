package avltree

// W.I.P.

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

type avlNode struct {
	Left    *avlNode
	Right   *avlNode
	Parent  *avlNode
	Element int
}

type AvlTree struct {
	Root *avlNode
	size int
}

func (n *avlNode) isLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *avlNode) height() int {
	if n.isLeaf() {
		return 0
	}
	return max(n.Left.height(),
		n.Right.height())
}

func (n *avlNode) balance() int {
	return abs(n.Left.height() - n.Right.height())
}

func (n *avlNode) binarySearch(target int) *avlNode {
	if n.isLeaf() || n.Element == target {
		return n
	} else if target > n.Element {
		return n.Right.binarySearch(target)
	} else {
		return n.Left.binarySearch(target)
	}
}

func (n *avlNode) expandLeaf(element int) {
	if !n.isLeaf() {
		return
	}
	n.Left = &avlNode{Parent: n}
	n.Right = &avlNode{Parent: n}
	n.Element = element
}

func (t *AvlTree) insert(newElement int) bool {
	if t.size == 0 {
		t.Root.expandLeaf(newElement)
		t.size++
		return true
	}
	result := t.Root.binarySearch(newElement)
	if !result.isLeaf() {
		return false
	}
	result.expandLeaf(newElement)
	return true
}

func (n *avlNode) fixUp() {

}

func rotateLeft(child *avlNode, parent *avlNode) {
	parent.Right = child.Left
	child.Left = parent
	child.Parent = parent.Parent
	parent.Parent = child
}

func rotateRight(child *avlNode, parent *avlNode) {
	//parent.Right = child.Left
	//child.Left = parent
	//child.Parent = parent.Parent
	//parent.Parent = child
}
