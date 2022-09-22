package my_list

import (
	"fmt"
	"testing"
)

func TestCLink(t *testing.T) {

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

func Test_DealCards(t *testing.T) {
	my_list := new(CLinkList)
	my_list.DealCards()
}

func Test_JosephRing(t *testing.T) {
	my_list := new(CLinkList)
	my_list.JosephRing()
}
