package my_list

//
//import (
//	"fmt"
//)
//
////链式存储
//type LinkList struct {
//	Head   *Node
//	Length int
//}
//
//type Node struct {
//	Data ElemType
//	Next *Node
//}
//
//func NewLinkNode(value ElemType) *Node {
//	return &Node{
//		value,
//		nil,
//	}
//}
//
////func (this  *Node) GetNext() *Node {
////	//return this.Next
////}
//
////初始化列表
//func (l *LinkList) InitList() {
//	l.Length = 0
//	l.Head = NewLinkNode(EmptyElement)
//}
//
////清空列表
//func (l *LinkList) ClearList() {
//	p := new(Node)
//	q := l.Head
//	for {
//		if q != nil {
//			p = q
//			q = p.Next
//			p = nil
//		} else {
//			break
//		}
//	}
//	l.Head.Next = nil
//	l.Length = 0
//}
//
////判断是否为空
//func (l *LinkList) ListEmpty() bool {
//	if l.Length > 0 {
//		return false
//	}
//	return true
//}
//
////获取长度
//func (l *LinkList) ListLength() int {
//	return l.Length
//}
//
////增
//func (l *LinkList) AddItem(index int, value ElemType) bool {
//
//	if index < 0 || index > l.Length {
//		return false
//	}
//
//	index = index + 1 //头结点位置
//
//	pItem := l.Head
//
//	for j := 1; j < index; j++ {
//		pItem = pItem.Next
//	}
//
//	newItem := NewLinkNode(value)
//	pItem.Next = newItem
//
//	l.Length++
//
//	return true
//}
//
////删
//func (l *LinkList) DelItem(index int) bool {
//	if l.Length == 0 {
//		return false
//	}
//	if index < 0 || index > l.Length {
//		return false
//	}
//
//	index = index + 1 //头结点位置
//
//	preItem := l.Head
//	curItem := l.Head.Next
//
//	delItemDate := l.GetItem(index)
//
//	for nil != curItem {
//		if curItem.Data == delItemDate {
//			break
//		}
//		preItem = curItem
//		curItem = curItem.Next
//	}
//
//	//要删除的节点不存在
//	if curItem == nil {
//		return false
//	}
//
//	preItem.Next = curItem.Next
//	l.Length--
//
//	return true
//}
//
//func getNodeByIndex(l *LinkList, index int) *Node {
//	if index > l.Length || index < 0 {
//		return nil
//	}
//	preItem := l.Head.Next
//	for i := 0; i < index; i++ {
//		preItem = preItem.Next
//	}
//	return preItem
//}
//
////改
//func (l *LinkList) SetItem(index int, value ElemType) bool {
//	if index > l.Length || index < 0 {
//		return false
//	}
//	preItem := getNodeByIndex(l, index)
//	if preItem == nil {
//		return false
//	}
//	preItem.Data = value
//	return true
//}
//
////查
//func (l *LinkList) GetElem(index int) ElemType {
//	if index < 0 || index > l.Length {
//		return EmptyElement
//	}
//	if l.Length == 0 {
//		return EmptyElement
//	}
//
//	preItem := getNodeByIndex(l, index)
//
//	return preItem.Data
//}
//
////输出
//func (l *LinkList) Echo() {
//	curItem := l.Head.Next
//	for i := 0; i < l.Length; i++ {
//		fmt.Println(curItem.Data, i)
//		curItem = curItem.Next
//	}
//	fmt.Println("end")
//}
//
//func (l *LinkList) Test() {
//	my_list := new(LinkList)
//	my_list.InitList()
//	for i := 0; i < 10; i++ {
//		my_list.AddItem(i, ElemType(i*i))
//	}
//
//	my_list.Echo()
//	my_list.DelItem(2)
//	res := my_list.GetLength()
//	fmt.Println(res)
//	my_list.Echo()
//	my_list.SetItem(0, 2)
//	value := my_list.GetItem(0)
//	fmt.Println(value)
//
//	my_list.ClearList()
//	res1 := my_list.CheckEmpty()
//	println(res1)
//	my_list.Echo()
//}
