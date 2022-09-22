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

func (l *LinkList) InitTestList() {
	l.InitList()

	for i := 1; i <= 11; i++ {
		l.ListInsert(i, ElemType(i*i+1))
		l.Echo()
	}
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

/**
题目
快速找到未知长度的单链表的中间节点
解决思路
（一）使用一个指针，先索引一遍获取总长度，再取长度一半去循环获取到中间值
算法复杂度：O(L)+O(L/2) = O(3L/2)
 （二）使用两个指针，快指针和慢指针，快指针一次向前走2格，慢指针一次走一格，当快指针走完全程，慢指针正好走在中间
算法复杂度：O(L/2)
*/

func (l *LinkList) FindMidEle(e *ElemType) int {
	//search是快指针，middle是慢指针
	search := l.Head
	middle := l.Head

	i := 0
	for search != nil {
		if search.Next != nil {
			//快指针是慢指针的两倍速度
			search = search.Next.Next
			middle = middle.Next
		} else { //是针对奇数个元素，再次进行操作，是之能够退出
			search = search.Next
			middle = middle.Next //奇数需要再次向后取一位，才会到达中间
		}
		i++
	}
	*e = middle.Data
	return i
}

/**
约瑟夫环
19个人报数，1-3，当谁报数为3，谁就淘汰。现在获取他们淘汰的顺序
*/
func (l *CLinkList) JosephRing() {
	my_list := new(CLinkList)
	my_list.InitList()

	for i := 1; i <= 19; i++ {
		my_list.ListInsert(i, ElemType(i))
		//my_list.Echo()
	}
	var e ElemType

	var s1 []ElemType //创建nil切换

	cur := my_list.Head.Next
	for i := 1; my_list.ListEmpty() != true; i++ {
		index := my_list.LocateElem(cur.Data)
		fmt.Println("第几个:", index, "数值:", cur.Data, "喊,", i)
		if i%3 == 0 {
			my_list.ListDelete(index, &e)
			fmt.Println("淘汰:", e)
			s1 = append(s1, e)
			my_list.Echo()
			i = 0
		}
		cur = cur.Next
		if cur.Data == 0 { //跳过头结点
			cur = cur.Next
		}
	}

	fmt.Println("淘汰顺序:", s1)

	fmt.Println("结束")
}
