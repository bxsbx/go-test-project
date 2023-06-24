package util

import "errors"

type Stack[T any] struct {
	top  int // 栈顶元素位置
	list []T
}

// 创建栈，可以初始化容量
func NewStack[T any](size ...int) Stack[T] {
	stack := Stack[T]{}
	if len(size) > 0 && size[0] > 0 {
		stack.list = make([]T, size[0])
	} else {
		stack.list = make([]T, 10)
	}
	stack.top = -1
	return stack
}

// 入栈
func (s *Stack[T]) Push(obj T) {
	s.top++
	if len(s.list) > s.top {
		s.list[s.top] = obj
	} else {
		s.list = append(s.list, obj)
	}
}

// 出栈
func (s *Stack[T]) Pull() (T, error) {
	if s.top < 0 {
		return s.list[0], errors.New("栈中已没有元素,无法出栈")
	}
	t := s.list[s.top]
	s.top--
	return t, nil
}

// 获取栈顶元素
func (s Stack[T]) Top() (T, error) {
	if s.top < 0 {
		return s.list[0], errors.New("栈中已没有元素")
	}
	return s.list[s.top], nil
}

// 获取栈元素个数
func (s Stack[T]) Size() int {
	return s.top + 1
}

// 缩减栈的容量
func (s *Stack[T]) CutCapSize() {
	if len(s.list) > 20 && len(s.list) > (s.top+1)*2 {
		list := make([]T, 0)
		s.list = append(list, s.list[0:s.top+1]...)
	}
}

// 返回栈的所有元素
func (s Stack[T]) GetAllElem() []T {
	if s.top < 0 {
		return []T{}
	}
	return s.list[:s.top+1]
}

// 栈定元素是否存在
func (s Stack[T]) IsExistElem() bool {
	if s.top < 0 {
		return false
	}
	return true
}
