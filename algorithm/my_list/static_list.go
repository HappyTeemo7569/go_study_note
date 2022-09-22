package my_list

import (
	"fmt"
)

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
	//初始化使还没有数据，所以这时头结点指针指向0，代表空
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

//获取指定位置的元素，返回在指针元素中
func (l *StaticList) GetElem(index int, e *ElemType) bool {

	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}
	if index < 1 || index > l.Length {
		fmt.Println("获取失败，位置错误")
		return false
	}

	//获取头结点
	index_first := l.Element[MAXSIZE-1].Cur

	for j := 1; index_first != 0; j++ {
		if j == index {
			*e = l.Element[index_first].Data
			break
		}
		index_first = l.Element[index_first].Cur
	}

	return true
}

//查找元素在线性表中的位置
func (l *StaticList) LocateElem(value ElemType) int {

	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return 0
	}

	//获取头结点
	index_first := l.Element[MAXSIZE-1].Cur

	j := 1
	for ; j < MAXSIZE; j++ {
		if l.Element[index_first].Data == value {
			break
		}
		index_first = l.Element[index_first].Cur
	}

	return j
}

//获取备份链表的第一个可用下标
func (l *StaticList) Malloc_SLL() int {

	//获取当前数组的第一个元素中存放的可用下标
	i := l.Element[0].Cur

	//若是下标可用，若是指向0，数组空间全部用尽
	if l.Element[0].Cur != 0 {
		l.Element[0].Cur = l.Element[i].Cur
	}

	return i
}

//增
func (l *StaticList) ListInsert(index int, value ElemType) bool {

	if l.Length == MAXSIZE { //满了
		fmt.Println("插入失败，队列已满")
		return false
	}
	if index < 1 || index > l.Length+1 {
		fmt.Println("插入失败，位置错误")
		return false
	}

	//获取空闲分量的下标
	free_index := l.Malloc_SLL()

	k := MAXSIZE - 1 //k首先是最后一个元素的下标
	if free_index != 0 {
		l.Element[free_index].Data = value
		for i := 1; i < index-1; i++ { //找到第i个元素的之前的位置
			k = l.Element[k].Cur
		}
		//把第i个元素之前的cur赋值给新的元素的cur
		l.Element[free_index].Cur = l.Element[k].Cur
		l.Element[k].Cur = free_index

		l.Length++
		return true
	}

	return false
}

//将回退首数组的数据，指向我们删除的那个空间，不会造成内存碎片
func (l *StaticList) Free_SLL(k int) {

	//将我们删除的这个元素游标指向原来的下一个空闲分量，方便一会回到原来的位置
	l.Element[k].Cur = l.Element[0].Cur

	//将我们删除的那个元素放入到0下标的游标中，在下一次插入时会被调用，存放数据，节约空间
	l.Element[0].Cur = k
}

//删
func (l *StaticList) ListDelete(index int, e *ElemType) bool {
	if l.Length == 0 {
		fmt.Println("获取失败，队列为空")
		return false
	}
	if index < 1 || index > l.Length {
		fmt.Println("获取失败，位置错误")
		return false
	}

	first_index := MAXSIZE - 1 //头结点位置

	//注意下标index开始和j的关系，我们要删除第三个，j会循环两次，
	// 我们若是从index = space[MAXSIZE-1].cur（是指向第一个）开始，
	// 循环两次会直接指向我们要删除的那个，而不是我们想要的前一个
	j := 0

	//找到我们要删除的元素的前一个i-1元素的下标
	for j = 1; j < index; j++ {
		first_index = l.Element[first_index].Cur
	}
	//用来保存原来删除的那个元素下标
	j = l.Element[first_index].Cur
	*e = l.Element[j].Data
	//使前一个元素游标指向要删除元素下一个元素的下标，跳过中间这个要删除的
	l.Element[first_index].Cur = l.Element[j].Cur
	//将我们删除的那个元素下标放入到备用链表中
	l.Free_SLL(j)

	l.Length--
	return true
}

//输出
func (l *StaticList) Echo() {
	start := l.Element[MAXSIZE-1].Cur
	index := start
	for index != 0 {
		fmt.Print(l.Element[index].Data, " ")
		index = l.Element[index].Cur
	}
	fmt.Println()
}
