package my_stack

import "fmt"

//栈因为只是栈顶来做插入和删除操作，所以比较好的方法就是将栈顶放在单链表的头部，栈顶指针和单链表的头指针合二为一

//顺序存储
type LinkStack struct {
	//Head   *Node
	Top    *Node //注意：由于栈顶是在开始结点，会一直变化，我们不需要设置头结点  其实就是一直头插
	Length int
}
type Node struct {
	Data ElemType //数据
	Next *Node    //下一个节点
}

//初始化操作，建立一个空栈
func (s *LinkStack) InitStack() {
	s.Top = nil
	s.Length = 0
}

//将栈清空,将栈顶指针移动到栈底即可，容量大小不要修改，数据不需要清空，数据入栈会覆盖
func (s *LinkStack) ClearStack() {
	s.Top = nil
	s.Length = 0
}

//判断是否为空
func (s *LinkStack) StackEmpty() bool {
	if s.Length == 0 {
		return true
	}
	return false
}

//获取栈的长度
func (s *LinkStack) StackLength() int {
	return s.Length
}

//获取栈顶
func (s *LinkStack) GetTop(e *ElemType) bool {
	if s.StackEmpty() {
		return false
	}
	*e = s.Top.Data
	return true
}

//进栈
func (s *LinkStack) Push(value ElemType) bool {
	if s.Length >= MAXSIZE {
		return false
	}

	n := new(Node)
	n.Next = s.Top
	n.Data = value
	s.Top = n

	s.Length++

	return true
}

//出站
func (s *LinkStack) Pop(e *ElemType) bool {
	if s.StackEmpty() {
		return false
	}

	tmp := s.Top
	s.Top = tmp.Next
	*e = tmp.Data

	s.Length--
	return true
}

//销毁栈
func (s *LinkStack) DestroyStack() bool {
	return true
}

func (s *LinkStack) Echo() {
	var e ElemType
	for s.Length > 0 {
		s.Pop(&e)
		fmt.Print(e, " ")
	}
	fmt.Println()
}

func (s *LinkStack) Test() {
	fmt.Println("测试开始")

	my_list := new(LinkStack)
	my_list.InitStack()

	for i := 1; i <= 10; i++ {
		my_list.Push(ElemType(i*i + 1))
	}

	my_list.Echo()

	var e ElemType

	for i := 1; i <= 10; i++ {
		my_list.Push(ElemType(i*i + 1))
	}
	my_list.Pop(&e)
	fmt.Println("删除头元素:", e)
	my_list.Echo()

	for i := 1; i <= 10; i++ {
		my_list.Push(ElemType(i*i + 1))
	}
	my_list.GetTop(&e)
	fmt.Println("获取栈顶:", e)

	for i := 1; i <= 10; i++ {
		my_list.Push(ElemType(i*i + 1))
	}
	fmt.Println("开始清空")
	my_list.ClearStack()
	if my_list.StackEmpty() {
		fmt.Println("已清空")
		my_list.Echo()
	}

	fmt.Println("测试完成")
}
