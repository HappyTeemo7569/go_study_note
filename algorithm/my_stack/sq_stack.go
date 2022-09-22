package my_stack

import "fmt"

//go不能做指针运算，所以不能像C一样用指针位移操作栈
//改为用切片实现

//顺序存储
type SqStack struct {
	Top     int
	Element [MAXSIZE]ElemType //动态扩容
}

//初始化操作，建立一个空栈
func (s *SqStack) InitStack() {
	s.Top = -1
	for i := 0; i < MAXSIZE; i++ {
		s.Element[i] = EmptyElement
	}
}

//将栈清空,将栈顶指针移动到栈底即可，容量大小不要修改，数据不需要清空，数据入栈会覆盖
func (s *SqStack) ClearStack() {
	s.Top = -1
	for i := 0; i < MAXSIZE; i++ {
		s.Element[i] = EmptyElement
	}
}

//判断是否为空
func (s *SqStack) StackEmpty() bool {
	if -1 == s.Top {
		return true
	}
	return false
}

//获取栈的长度
func (s *SqStack) StackLength() int {
	return s.Top + 1
}

//获取栈顶
func (s *SqStack) GetTop(e *ElemType) bool {
	if s.StackEmpty() {
		return false
	}
	*e = s.Element[s.Top]
	return true
}

//进栈
func (s *SqStack) Push(value ElemType) bool {
	if s.Top+1 >= MAXSIZE {
		return false
	}
	s.Top++
	s.Element[s.Top] = value
	return true
}

//出站
func (s *SqStack) Pop(e *ElemType) bool {
	if s.StackEmpty() {
		return false
	}
	*e = s.Element[s.Top]
	s.Top--
	return true
}

//销毁栈
func (s *SqStack) DestroyStack() bool {
	return true
}

func (s *SqStack) Echo() {
	fmt.Println(s.Element)
}
