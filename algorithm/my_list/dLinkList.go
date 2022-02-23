package my_list

import (
	"fmt"
	"math"
	"my_go/dataStructure/list"
)

//双向循环链表
type DLinkList struct {
	Head   *DNode
	Length int
}

type DNode struct {
	Data  list.ElemType
	Next  *DNode //直接后继指针
	Prior *DNode //直接前驱指针
}

//初始化列表
func (l *DLinkList) InitList() {
	l.Head = new(DNode)
	l.Head.Next = l.Head  //指向自己
	l.Head.Prior = l.Head //指向自己
	l.Length = 0
}

//清空列表（不会清除头结点）
func (l *DLinkList) ClearList() {
	l.Head.Next = l.Head  //指向自己
	l.Head.Prior = l.Head //指向自己
	l.Length = 0
}

//判断是否为空
func (l *DLinkList) ListEmpty() bool {
	if l.Length == 0 {
		return true
	}
	return false
}

//获取长度
func (l *DLinkList) ListLength() int {
	return l.Length
}

//查  index可以正 可以负 负表示往前找
func (l *DLinkList) GetElem(index int, e *list.ElemType) bool {
	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}

	i := int(math.Abs(float64(index)))

	if index > list.MAXSIZE {
		index = index + 1 //跳过头结点
	} else if index < 0 {
		index = index - 1 //跳过头结点
	}

	j := 1
	q := l.Head

	if index > 0 {
		q = q.Next
	} else {
		q = q.Prior
	}
	for q != nil && j < i {
		if index > 0 {
			q = q.Next
		} else {
			q = q.Prior
		}

		j++
	}

	*e = q.Data
	return true
}

//按照元素进行查找，获取索引
func (l *DLinkList) LocateElem(value list.ElemType) int {

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

	if j >= list.MAXSIZE {
		return 0
	}

	return j
}

//按照索引进行插入数据  可以双向插入
func (l *DLinkList) ListInsert(index int, value list.ElemType) bool {

	if l.Length == list.MAXSIZE { //满了
		fmt.Println("插入失败，队列已满")
		return false
	}
	//if index < 1 || index > l.Length+1 {
	//	fmt.Println("插入失败，位置错误")
	//	return false
	//}

	i := int(math.Abs(float64(index)))

	if index > list.MAXSIZE {
		index = index + 1 //跳过头结点
	} else if index < 0 {
		index = index - 1 //跳过头结点
	}

	j := 1
	cur := l.Head

	if index > 0 {
		cur = cur.Next
	} else {
		cur = cur.Prior
	}
	for cur != nil && j < i {
		if index > 0 {
			cur = cur.Next
		} else {
			cur = cur.Prior
		}
		j++
	}

	//新建节点，加入链表
	n := new(DNode)
	if index > 0 {
		n.Next = cur
		n.Prior = cur.Prior
		cur.Prior.Next = n
		cur.Prior = n
	} else {
		n.Next = cur.Next
		n.Prior = cur
		cur.Next.Prior = n
		cur.Next = n
	}
	n.Data = value

	l.Length++

	return true
}

//删
func (l *DLinkList) ListDelete(index int, e *list.ElemType) bool {
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
	tmp.Next.Prior = front.Prior
	//free(tmp)

	l.Length--

	return true
}

//输出
func (l *DLinkList) Echo() {
	//遍历的写法
	curItem := l.Head.Next
	for i := 0; i < l.Length; i++ {
		fmt.Print(curItem.Data, " ")
		curItem = curItem.Next
	}
	fmt.Println()
}

func (l *DLinkList) Test() {
	fmt.Println("测试开始")

	my_list := new(DLinkList)
	my_list.InitList()

	for i := 1; i <= 10; i++ {
		my_list.ListInsert(i, list.ElemType(i*i+1))
		my_list.Echo()
	}

	fmt.Println("第5个这里插入256")
	my_list.ListInsert(5, 256)
	my_list.Echo()

	fmt.Println("第-5个这里插入256")
	my_list.ListInsert(-5, 256)
	my_list.Echo()

	my_list.ListInsert(199, 99)

	var e list.ElemType

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

	fmt.Println("测试完成")
}
