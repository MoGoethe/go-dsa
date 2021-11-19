package dst

// 二叉搜索树
// type LinkList str uct{
// 	items []interface{}
// }

type BinarySearchTreeNode struct {
	Key   int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func (t *BinarySearchTree) Insert(k int) {
	n := t.root
	newNode := &BinarySearchTreeNode{Key: k}
	if n == nil {
		t.root = newNode
	} else {
		t.insertNode(n, newNode)
	}
}
func (t *BinarySearchTree) insertNode(n1 *BinarySearchTreeNode, n2 *BinarySearchTreeNode) {
	if n1.Key > n2.Key {
		if n1.Left == nil {
			n1.Left = n2
		} else {
			t.insertNode(n1.Left, n2)
		}
	} else {
		if n1.Right == nil {
			n1.Right = n2
		} else {
			t.insertNode(n1.Right, n2)
		}
	}
}
func (t *BinarySearchTree) Has(key int) bool {
	return t.has(t.root, key)
}
func (t *BinarySearchTree) has(n1 *BinarySearchTreeNode, key int) bool {
	if n1 == nil {
		return false
	}
	if n1.Key > key {
		return t.has(n1.Left, key)
	}
	if n1.Key < key {
		return t.has(n1.Right, key)
	}
	return true
}
func (t *BinarySearchTree) Remove(key int) {
	t.remove(t.root, key)
}
func (t *BinarySearchTree) remove(n *BinarySearchTreeNode, key int) bool {
	if n == nil {
		return false
	}
	if n.Key > key {
		return t.remove(n.Left, key)
	}
	if n.Key < key {
		return t.remove(n.Right, key)
	}
	// 移除节点
	// 左右节点均为空，直接移除
	// 只存在左节点
	// 只存在右节点
	// 左右节点都存在
	return true
}
func (t *BinarySearchTree) Min() {}
func (t *BinarySearchTree) Max() {}

// 前序遍历
func (t *BinarySearchTree) PreOrderTraverse(cb func(*BinarySearchTreeNode)) {
	t.preOrderTraverseNode(t.root, cb)
}
func (t *BinarySearchTree) preOrderTraverseNode(n *BinarySearchTreeNode, cb func(*BinarySearchTreeNode)) {
	if n != nil {
		cb(n)
		t.preOrderTraverseNode(n.Left, cb)
		t.preOrderTraverseNode(n.Right, cb)
	}
}

// 中序遍历
func (t *BinarySearchTree) InOrderTraverse(cb func(*BinarySearchTreeNode)) {
	t.inOrderTraverseNode(t.root, cb)
}
func (t *BinarySearchTree) inOrderTraverseNode(n *BinarySearchTreeNode, cb func(*BinarySearchTreeNode)) {
	if n != nil {
		t.inOrderTraverseNode(n.Left, cb)
		cb(n)
		t.inOrderTraverseNode(n.Right, cb)
	}
}

// 后序遍历
func (t *BinarySearchTree) PostOrderTraverse(cb func(*BinarySearchTreeNode)) {
	t.postOrderTraverseNode(t.root, cb)
}
func (t *BinarySearchTree) postOrderTraverseNode(n *BinarySearchTreeNode, cb func(*BinarySearchTreeNode)) {
	if n != nil {
		t.postOrderTraverseNode(n.Left, cb)
		t.postOrderTraverseNode(n.Right, cb)
		cb(n)
	}
}

// 深度优先遍历
func (t *BinarySearchTree) DepthFirstTraverse() {
}

// 广度优先遍历
func (t *BinarySearchTree) BreadthFirstTraverse() {
}
