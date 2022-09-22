package my_list

import (
	"fmt"
	"testing"
)

func (l *LinkList) Test() {
	fmt.Println("测试开始")

	my_list := new(LinkList)
	my_list.InitList()

	for i := 1; i <= 10; i++ {
		my_list.ListInsert(i, ElemType(i*i+1))
		my_list.Echo()
	}

	fmt.Println("第5个这里插入256")
	my_list.ListInsert(5, 256)
	my_list.Echo()
	my_list.ListInsert(199, 99)

	var e ElemType

	my_list.ListDelete(1, &e)
	fmt.Println("删除头元素:", e)
	my_list.Echo()

	my_list.ListDelete(my_list.ListLength(), &e)
	fmt.Println("删除尾元素:", e)
	my_list.Echo()

	my_list.GetElem(6, &e)
	fmt.Println("获取第6个:", e)

	fmt.Println("256的位置:", my_list.LocateElem(256))

	fmt.Println("长度：", my_list.ListLength())

	fmt.Println("开始清空")
	my_list.ClearList()
	if my_list.ListEmpty() {
		fmt.Println("已清空")
		my_list.Echo()
	}

	//fmt.Println("准备合并")
	//
	//my_list_a := new(SqList)
	//my_list_b := new(SqList)
	//for i := 1; i <= 10; i++ {
	//	my_list_a.ListInsert(i, ElemType(2*i+1))
	//	my_list_b.ListInsert(i, ElemType(3*i+1))
	//}
	//
	//my_list_a.Echo()
	//my_list_b.Echo()
	//
	//fmt.Println("合并后")
	//
	//my_list_a.UnionL(my_list_b)
	//my_list_a.Echo()

	fmt.Println("测试完成")
}

func Test1(t *testing.T) {
	my_list := new(LinkList)
	my_list.InitTestList()
	var e ElemType
	index := my_list.FindMidEle(&e)
	fmt.Println("中间值:", e, "位置", index)
}
