package my_list

import (
	"fmt"
)

type CLinkList struct {
	Head   *Node
	Length int
}

//初始化列表
func (l *CLinkList) InitList() {
	l.Head = new(Node)
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
	n := new(Node)
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

/**
魔术师发牌问题
一共13张黑牌1-13，预先排好顺序，牌面朝下，开始数数，
数1翻开第一张牌为1取出，
数1,2，将喊到1的牌放在末尾，喊到2的牌翻开为2取出
喊1,2,3，将喊到1和2的牌放到末尾，将喊到3的牌翻开取出
......
喊到13，直到将所有牌取出。
正好翻开顺序为1-2-....-13
我们需要找到他排列的预先顺序19个人报数，1-3，当谁报数为3，谁就淘汰。现在获取他们淘汰的顺序

注意：翻开一张少一张
*/
func (l *CLinkList) DealCards() {

	card_num := 1 //卡牌编号

	card_count := 13 //卡牌数量

	my_list := new(CLinkList)
	my_list.InitList()
	for i := 1; i <= card_count; i++ {
		my_list.ListInsert(i, 100)
	}
	my_list.Echo()

	//start := my_list.Head.Next
	start := my_list.Head
	for card_num <= card_count {

		for j := 1; j <= card_num; j++ {
			start = start.Next
			for start.Data != 100 { //跳过头结点 找空白点
				start = start.Next
			}
		}

		start.Data = ElemType(card_num)
		fmt.Println(card_num)
		my_list.Echo()

		if card_num == card_count {
			break
		} //注意这个判断要在while前面，不然会造成死循环

		for start.Next.Data != 100 {
			start = start.Next
		}
		card_num++
	}

	//1 8 2 5 10 3 12 11 9 4 7 6 13

	fmt.Println("开始输出")
	start = my_list.Head
	card_num = 1
	var e ElemType
	for i := 0; i < card_count; i++ {

		for j := 1; j <= card_num; j++ {
			start = start.Next
			if start.Data == 0 { //跳过头结点
				start = start.Next
			}
		}

		fmt.Println(start.Data)
		my_list.ListDelete(my_list.LocateElem(start.Data), &e)
		card_num++
	}
}
