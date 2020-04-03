package list

import "fmt"

//静态链表
type StaticList struct {
	Element [MAXSIZE]StaticNode
	Length  int
}

type StaticNode struct {
	Data ElemType
	Cur  int
}

//静态链表理论上是要用参数传递的。

//初始化列表
func (l *StaticList) InitList() {
	l.Length = 0
	for i := 0; i < MAXSIZE; i++ {
		l.Element[i] = StaticNode{EmptyElement, i + 1}
	}
	l.Element[MAXSIZE-1].Cur = 0
}

//清空列表 就是将首尾两个数组设置为原始的1和0
func (l *StaticList) ClearList() {
	l.Element[0].Cur = 1
	l.Element[MAXSIZE-1].Cur = 0
}

//判断是否为空
func (l *StaticList) ListEmpty() bool {
	if l.Length > 0 {
		return false
	}
	return true
}

//获取长度
func (l *StaticList) ListLength() int {
	return l.Length
}

//增
func (l *StaticList) AddItem(index int, value ElemType) bool {
	if l.Length == MAXSIZE {
		return false
	}
	if index < 0 || index > l.Length {
		return false
	}

	for k := l.Length; k > index; k-- {
		l.Element[k] = l.Element[k-1]
	}
	l.Element[index] = value
	l.Length++

	return true
}

//删
func (l *StaticList) DelItem(index int) bool {
	if l.Length == 0 {
		return false
	}
	if index < 0 || index > l.Length {
		return false
	}

	index = index + 1 //头结点位置

	preItem := l.Head
	curItem := l.Head.Next

	delItemDate := l.GetItem(index)

	for nil != curItem {
		if curItem.Data == delItemDate {
			break
		}
		preItem = curItem
		curItem = curItem.Next
	}

	//要删除的节点不存在
	if curItem == nil {
		return false
	}

	preItem.Next = curItem.Next
	l.Length--

	return true
}

//改
func (l *StaticList) SetItem(index int, value ElemType) bool {
	if index > l.Length || index < 0 {
		return false
	}
	preItem := getNodeByIndex(l, index)
	if preItem == nil {
		return false
	}
	preItem.Data = value
	return true
}

//查
func (l *StaticList) GetElem(index int) ElemType {
	if index < 0 || index > l.Length {
		return EmptyElement
	}
	if l.Length == 0 {
		return EmptyElement
	}

	preItem := getNodeByIndex(l, index)

	return preItem.Data
}

//输出
func (l *StaticList) Echo() {
	curItem := l.Head.Next
	for i := 0; i < l.Length; i++ {
		fmt.Println(curItem.Data, i)
		curItem = curItem.Next
	}
	fmt.Println("end")
}

func (l *StaticList) Test() {
	my_list := new(LinkList)
	my_list.InitList()
	for i := 0; i < 10; i++ {
		my_list.AddItem(i, ElemType(i*i))
	}

	my_list.Echo()
	my_list.DelItem(2)
	res := my_list.GetLength()
	fmt.Println(res)
	my_list.Echo()
	my_list.SetItem(0, 2)
	value := my_list.GetItem(0)
	fmt.Println(value)

	my_list.ClearList()
	res1 := my_list.CheckEmpty()
	println(res1)
	my_list.Echo()
}
