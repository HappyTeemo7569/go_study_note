package my_list

import (
	"fmt"
)

//链式存储
//可以用头指针用来存储长度 也就是第0个
//但是这个有个弊端就是当存储的不是int类型时会有问题
//所有还是添加一个字段表示长度
//为了保持从1开始，头指针还是不放东西
type LinkList struct {
	Head   *Node
	Length int
}

type Node struct {
	Data ElemType
	Next *Node
}

//初始化列表
func (l *LinkList) InitList() {
	l.Head = new(Node)
	//l.Head.Next = new(Node)
	l.Length = 0
}

//清空列表（不会清除头结点）
func (l *LinkList) ClearList() {
	p := new(Node)
	q := l.Head //q指向第一个结点

	//释放内存，其实go可以不需要这个循环
	for q != nil {
		p = q
		q = p.Next
		p = nil
	}
	l.Head.Next = nil
	l.Length = 0
}

//判断是否为空
func (l *LinkList) ListEmpty() bool {
	if l.Length == 0 {
		return true
	}
	return false
}

//获取长度
func (l *LinkList) ListLength() int {
	return l.Length
}

//查
func (l *LinkList) GetElem(index int, e *ElemType) bool {
	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}
	if index < 1 || index > l.Length {
		fmt.Println("获取失败，位置错误")
		return false
	}

	j := 1
	q := l.Head.Next
	for q != nil && j < index {
		q = q.Next
		j++
	}

	//有了这一步其实开头的判断可以去掉
	if q == nil || j > index {
		return false
	}

	*e = q.Data
	return true
}

//按照元素进行查找，获取索引
func (l *LinkList) LocateElem(value ElemType) int {

	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return 0
	}

	j := 0
	q := l.Head.Next
	for q != nil {
		j++
		if q.Data == value {
			break
		}
		q = q.Next
	}

	if j >= MAXSIZE {
		return 0
	}

	return j
}

//按照索引进行插入数据
func (l *LinkList) ListInsert(index int, value ElemType) bool {

	if l.Length == MAXSIZE { //满了
		fmt.Println("插入失败，队列已满")
		return false
	}
	if index < 1 || index > l.Length+1 {
		fmt.Println("插入失败，位置错误")
		return false
	}

	front := l.Head

	//找到插入位置的前驱
	for j := 1; j < index; j++ {
		front = front.Next
	}

	//新建节点，加入链表
	n := new(Node)
	n.Next = front.Next
	n.Data = value
	front.Next = n

	l.Length++

	return true
}

//删
func (l *LinkList) ListDelete(index int, e *ElemType) bool {
	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}
	if index < 1 || index > l.Length {
		fmt.Println("获取失败，位置错误")
		return false
	}

	j := 1
	front := l.Head

	//找到索引的直接前驱
	for front.Next != nil && j < index {
		front = front.Next
		j++
	}

	if front.Next == nil || j > index {
		return false
	}

	//开始删除
	tmp := front.Next     //记录要删除的
	*e = tmp.Data         //返回删除节点的数据
	front.Next = tmp.Next //前驱节点直接指向后继节点，就跳过了要删除的节点
	//free(tmp)

	l.Length--

	return true
}

//输出
func (l *LinkList) Echo() {
	//遍历的写法
	curItem := l.Head.Next
	for i := 0; i < l.Length; i++ {
		fmt.Print(curItem.Data, " ")
		curItem = curItem.Next
	}
	fmt.Println()
}

func (l *LinkList) Test() {
	fmt.Println("测试开始")

	my_list := new(LinkList)
	my_list.InitList()

	for i := 1; i <= 10; i++ {
		my_list.ListInsert(i, ElemType(i*i+1))
		my_list.Echo()
	}

	fmt.Println("第5个这里插入256")
	my_list.ListInsert(5, 256)
	my_list.Echo()
	my_list.ListInsert(199, 99)

	var e ElemType

	my_list.ListDelete(1, &e)
	fmt.Println("删除头元素:", e)
	my_list.Echo()

	my_list.ListDelete(my_list.ListLength(), &e)
	fmt.Println("删除尾元素:", e)
	my_list.Echo()

	my_list.GetElem(6, &e)
	fmt.Println("获取第6个:", e)

	fmt.Println("256的位置:", my_list.LocateElem(256))

	fmt.Println("长度：", my_list.ListLength())

	fmt.Println("开始清空")
	my_list.ClearList()
	if my_list.ListEmpty() {
		fmt.Println("已清空")
		my_list.Echo()
	}

	//fmt.Println("准备合并")
	//
	//my_list_a := new(SqList)
	//my_list_b := new(SqList)
	//for i := 1; i <= 10; i++ {
	//	my_list_a.ListInsert(i, ElemType(2*i+1))
	//	my_list_b.ListInsert(i, ElemType(3*i+1))
	//}
	//
	//my_list_a.Echo()
	//my_list_b.Echo()
	//
	//fmt.Println("合并后")
	//
	//my_list_a.UnionL(my_list_b)
	//my_list_a.Echo()

	fmt.Println("测试完成")
}
