package dst

// 队列

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []interface{}
}

// 入队
func (queue *Queue) EnQueue(element interface{}) (err error) {
	queue.items = append(queue.items, element)
	return nil
}

// 出队
func (queue *Queue) DeQueue() (ele interface{}, err error) {
	if len(queue.items) <= 0 {
		return nil, errors.New("队列为空")
	}
	last := queue.items[0]
	queue.items = queue.items[1:]
	return last, nil
}

// 获取队首
func (queue *Queue) Front() (ele interface{}, err error) {
	if len(queue.items) <= 0 {
		return nil, errors.New("队列为空")
	}
	return queue.items[0], nil
}

// 获取队尾
func (queue *Queue) Tail() (ele interface{}, err error) {
	if len(queue.items) <= 0 {
		return nil, errors.New("队列为空")
	}
	last := queue.items[len(queue.items)-1]
	return last, nil
}

// 判断是否为空
func (queue *Queue) IsEmpty() bool {
	return len(queue.items) == 0
}

// 清空
func (queue *Queue) Clear() {
	queue.items = queue.items[:0]
}

// 获取队列Size
func (queue *Queue) Size() int {
	return len(queue.items)
}

// 打印
func (queue *Queue) Print() {
	for i, v := range queue.items {
		fmt.Printf("第 %v 个元素是 %v, 它的类型是：%T\n", i, v, v)
	}
}
