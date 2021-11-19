package dst

import "errors"

// 栈
// type LinkList str uct{
// 	items []interface{}
// }

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("堆栈为空")
	}
	index := len(s.items) - 1
	result := s.items[index]
	s.items = s.items[:index]
	return result, nil
}
