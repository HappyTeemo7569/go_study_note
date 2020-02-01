package list

import "fmt"

type SqList struct {
	Element [MAXSIZE]ElemType
	Length  int
}

//初始化列表
func (l *SqList) InitList() {
	l.Length = 0

	for i := 0; i < MAXSIZE; i++ {
		l.Element[i] = EmptyElement
	}
}

//清空列表
func (l *SqList) ClearList() {
	l.Length = 0

	for i := 0; i < MAXSIZE; i++ {
		l.Element[i] = EmptyElement
	}
}

//判断是否为空
func (l *SqList) CheckEmpty() bool {
	if l.Length > 0 {
		return true
	}
	return false
}

//获取长度
func (l *SqList) GetLength() int {
	return l.Length
}

//增
func (l *SqList) AddItem(index int, value ElemType) bool {
	if l.Length == MAXSIZE {
		return false
	}
	if index < 0 || index > l.Length {
		return false
	}
	//开始进行插入操作，先判断元素插入位置，在将后面元素向后移动一位，进行插入
	for k := l.Length; k > index; k-- {
		l.Element[k] = l.Element[k-1]
	}
	l.Element[index] = value
	l.Length++
	return true
}

//删
func (l *SqList) DelItem(index int) bool {
	if l.Length == 0 {
		return false
	}
	if index < 0 || index > l.Length {
		return false
	}

	//删除元素，是将后面元素前移，再将最后的元素删除
	for k := index; k < l.Length-1; k++ {
		l.Element[k] = l.Element[k+1]
	}
	l.Element[l.Length-1] = EmptyElement
	l.Length--
	return true
}

//改
func (l *SqList) SetItem(index int, value ElemType) bool {
	if index > l.Length || index < 0 {
		return false
	}
	l.Element[index] = value
	return true
}

//查
func (l *SqList) GetItem(index int) ElemType {
	if index < 0 || index > l.Length {
		return EmptyElement
	}
	if l.Length == 0 {
		return EmptyElement
	}

	return l.Element[index]
}

//输出
func (l *SqList) Echo() {
	fmt.Println(l.Element)
}

func Run() {
	my_list := new(SqList)
	my_list.InitList()
	res := my_list.AddItem(0, 1)
	res = my_list.AddItem(1, 1)
	fmt.Println(res)
	my_list.Echo()
	res = my_list.DelItem(0)
	fmt.Println(res)
	my_list.Echo()
	my_list.SetItem(0, 2)
	value := my_list.GetItem(0)
	fmt.Println(value)
}
