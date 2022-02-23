package list

import (
	"fmt"
	"my_go/dataStructure/list/my_chain"
)

type CLinkList struct {
	Head   *my_chain.Node
	Length int
}

//初始化列表
func (l *CLinkList) InitList() {
	l.Head = new(my_chain.Node)
	l.Head.Next = l.Head //指向自己
	l.Length = 0
}

//清空列表（不会清除头结点）
func (l *CLinkList) ClearList() {
	//p := new(Node)
	//q := l.Head //q指向第一个结点

	//释放内存，其实go可以不需要这个循环
	//for q != nil {
	//	p = q
	//	q = p.Next
	//	p = nil
	//}
	l.Head.Next = l.Head
	l.Length = 0
}

//判断是否为空
func (l *CLinkList) ListEmpty() bool {
	if l.Length == 0 {
		return true
	}
	return false
}

//获取长度
func (l *CLinkList) ListLength() int {
	return l.Length
}

//查
func (l *CLinkList) GetElem(index int, e *ElemType) bool {
	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}
	//if index < 1 || index > l.Length {
	//	fmt.Println("获取失败，位置错误")
	//	return false
	//}

	if index > MAXSIZE {
		index = index + 1 //跳过头结点
	}
	j := 1
	q := l.Head.Next
	for q != nil && j < index {
		q = q.Next
		j++
	}

	//有了这一步其实开头的判断可以去掉
	//if q == nil || j > index {
	//	return false
	//}

	*e = q.Data
	return true
}

//按照元素进行查找，获取索引
func (l *CLinkList) LocateElem(value ElemType) int {

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
func (l *CLinkList) ListInsert(index int, value ElemType) bool {

	if l.Length == MAXSIZE { //满了
		fmt.Println("插入失败，队列已满")
		return false
	}
	if index < 1 {
		fmt.Println("插入失败，位置错误")
		return false
	}

	front := l.Head

	//找到插入位置的前驱
	for j := 1; j < index; j++ {
		front = front.Next
	}

	//新建节点，加入链表
	n := new(my_chain.Node)
	n.Next = front.Next
	n.Data = value
	front.Next = n

	l.Length++

	return true
}

//删
func (l *CLinkList) ListDelete(index int, e *ElemType) bool {
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
func (l *CLinkList) Echo() {
	//遍历的写法
	curItem := l.Head.Next
	for i := 0; i < l.Length; i++ {
		fmt.Print(curItem.Data, " ")
		curItem = curItem.Next
	}
	fmt.Println()
}

//实现两个链表的合并
func (la *CLinkList) Merge(lb *CLinkList) {

	//A的尾巴指向B的头，B的尾巴指向A的头
	a_last := la.Head.Next
	for j := 1; j < la.Length; j++ {
		a_last = a_last.Next
	}
	b_last := lb.Head.Next
	for j := 1; j < lb.Length; j++ {
		b_last = b_last.Next
	}

	//跳过B的头，也就是B的头干掉
	a_last.Next = lb.Head.Next
	b_last.Next = la.Head

	la.Length += lb.Length
}

func (l *CLinkList) Test() {
	fmt.Println("测试开始")

	my_list := new(CLinkList)
	my_list.InitList()

	for i := 1; i <= 19; i++ {
		my_list.ListInsert(i, ElemType(i*i+1))
		my_list.Echo()
	}

	fmt.Println("第5个这里插入256")
	my_list.ListInsert(5, 256)
	my_list.Echo()
	my_list.ListInsert(199, 99)

	var e ElemType

	my_list.GetElem(my_list.ListLength()+1, &e)
	fmt.Println("最后一个的下一个:", e)

	my_list.ListDelete(1, &e)
	fmt.Println("删除头元素:", e)
	my_list.Echo()

	my_list.ListDelete(my_list.ListLength(), &e)
	fmt.Println("删除尾元素:", e)
	my_list.Echo()

	my_list.GetElem(6, &e)
	fmt.Println("获取第6个:", e)

	//fmt.Println("最后一个的下一个:", e.Next)

	fmt.Println("256的位置:", my_list.LocateElem(256))

	fmt.Println("长度：", my_list.ListLength())

	fmt.Println("开始清空")
	my_list.ClearList()
	if my_list.ListEmpty() {
		fmt.Println("已清空")
		my_list.Echo()
	}

	fmt.Println("准备合并")

	my_list_a := new(CLinkList)
	my_list_a.InitList()
	my_list_b := new(CLinkList)
	my_list_b.InitList()

	for i := 1; i <= 10; i++ {
		my_list_a.ListInsert(i, ElemType(2*i+1))
		my_list_b.ListInsert(i, ElemType(3*i+1))
	}

	my_list_a.Echo()
	my_list_b.Echo()

	fmt.Println("合并后")

	my_list_a.Merge(my_list_b)
	my_list_a.Echo()
	my_list_a.GetElem(my_list_a.ListLength()+1, &e)
	fmt.Println("最后一个的下一个:", e)

	fmt.Println("测试完成")
}
