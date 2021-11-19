package dst

import "fmt"

// 链表
// type LinkList str uct{
// 	items []interface{}
// }

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

type LinkedList struct {
	head   *LinkedListNode
	length int
}

func (l *LinkedList) Prepend(n *LinkedListNode) {
	temp := l.head
	l.head = n
	l.head.Next = temp
	l.length++
}

func (l *LinkedList) Print() {
	n := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("第 %d 个数据是：%d ", i, n.Data)
		n = n.Next
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(data int) {
	if l.length == 0 {
		return
	}
	if l.head.Data == data {
		l.head = l.head.Next
		l.length--
		return
	}
	target := l.head
	for target.Next.Data != data {
		if target.Next.Next == nil {
			return
		}
		target = target.Next
	}
	target.Next = target.Next.Next
	l.length--
}
