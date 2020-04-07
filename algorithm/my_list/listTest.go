package my_list

import "fmt"

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

func (l *LinkList) Test1() {
	my_list := new(LinkList)
	my_list.InitList()

	for i := 1; i <= 11; i++ {
		my_list.ListInsert(i, ElemType(i*i+1))
		my_list.Echo()
	}
	var e ElemType
	index := my_list.FindMidEle(&e)
	fmt.Println("中间值:", e, "位置", index)
}
