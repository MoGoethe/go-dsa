package main

import (
	"datastructure/dst"
	"fmt"
)

func main() {
	fmt.Println("数据结构：")
	testHashTable()
	testTrie()
	testBinarySearchTree()
	testStack()
	testLinkedList()
	testMaxHeap()
	testQueue()
	testGraph()
}
func testGraph() {
	fmt.Println("图：")
	g := &dst.Graph{}
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(4)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 4)
	g.Print()
}

func testHashTable() {
	fmt.Println("哈希表：")
	h := dst.InitHashTable()
	h.Insert("hello")
	fmt.Println(h)
	h.Insert("world")
	h.Insert("world")
	fmt.Println(h.Search("hello"))
	fmt.Println(h.Delete("hello"))
	fmt.Println(h.Delete("hello"))
}

func testTrie() {
	fmt.Println("前缀树：")
	t := dst.InitTrie()
	t.Insert("hello")
	t.Insert("world")
	fmt.Println(t.Search("hello")) // true
}

func testBinarySearchTree() {
	fmt.Println("二叉搜索树：")
	tree := &dst.BinarySearchTree{}

	print := func(n *dst.BinarySearchTreeNode) {
		fmt.Print(n.Key, " ")
	}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(8)
	fmt.Println("前序遍历：")
	tree.PreOrderTraverse(print)
	fmt.Println()
	fmt.Println("中序遍历：")
	tree.InOrderTraverse(print)
	fmt.Println()
	fmt.Println("后序遍历：")
	tree.PostOrderTraverse(print)
	fmt.Println()
	fmt.Println(tree.Has(4), tree.Has(15))
}

func testStack() {
	fmt.Println("栈：")
	s := dst.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}

func testLinkedList() {
	fmt.Println("链表：")
	l := dst.LinkedList{}
	n1 := &dst.LinkedListNode{Data: 1, Next: nil}
	n2 := &dst.LinkedListNode{Data: 2, Next: nil}
	n3 := &dst.LinkedListNode{Data: 3, Next: nil}
	l.Prepend(n1)
	l.Prepend(n2)
	l.Prepend(n3)
	l.DeleteWithValue(4)
	l.Print()
	l.DeleteWithValue(3)
	l.Print()
}

func testMaxHeap() {
	fmt.Println("堆：")
	m := &dst.MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	fmt.Println(m)
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}

func testQueue() {
	fmt.Println("队列：")
	queue1 := &dst.Queue{}
	queue1.EnQueue(1)
	queue1.EnQueue(2)
	queue1.EnQueue(3)
	queue1.EnQueue(4)
	queue1.EnQueue(5)
	queue1.Print()
	fmt.Println(queue1)
	queue1.DeQueue()
	fmt.Println(queue1)
	fmt.Println("\n这个队列是空的吗？", queue1.IsEmpty())
}
