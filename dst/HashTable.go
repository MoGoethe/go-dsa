package dst

import "fmt"

// 哈希表
const HashTableSize int = 10

type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}
type HashTable struct {
	items [HashTableSize]*bucket
}

func InitHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.items {
		result.items[i] = &bucket{}
	}
	return result
}

func hash(key string) int {
	sum := 0
	for _, v := range key { // 对 string 进行循环
		sum += int(v)
	}
	return sum % HashTableSize
}

func (h *HashTable) Insert(key string) {
	if h.Search(key) == nil {
		hashCode := hash(key)
		h.items[hashCode].insertBucket(key)
	} else {
		fmt.Println(key, "已经存在了。")
	}
}
func (h *HashTable) Search(key string) *bucketNode {
	hashCode := hash(key)
	return h.items[hashCode].searchBucket(key)
}
func (h *HashTable) Delete(key string) bool {
	hashCode := hash(key)
	return h.items[hashCode].deleteBucket(key)
}

func (b *bucket) insertBucket(k string) {
	node := &bucketNode{key: k}
	node.next = b.head
	b.head = node
}
func (b *bucket) searchBucket(k string) *bucketNode {
	node := b.head
	for node != nil {
		if node.key == k {
			return node
		}
		node = node.next
	}
	return nil
}
func (b *bucket) deleteBucket(k string) bool {
	if b.head.key == k {
		b.head = b.head.next
		return true
	}
	node := b.head
	for node.next != nil {
		if node.next.key == k {
			node.next = node.next.next
			return true
		}
		node = node.next
	}
	fmt.Println(k, "不存在，无需删除。")
	return false
}
