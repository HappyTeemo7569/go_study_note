package my_list

import (
	"fmt"
	"testing"
)

func TestDLink(t *testing.T) {
	fmt.Println("测试开始")

	my_list := new(DLinkList)
	my_list.InitList()

	for i := 1; i <= 10; i++ {
		my_list.ListInsert(i, ElemType(i*i+1))
		my_list.Echo()
	}

	fmt.Println("第5个这里插入256")
	my_list.ListInsert(5, 256)
	my_list.Echo()

	fmt.Println("第-5个这里插入256")
	my_list.ListInsert(-5, 256)
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

	fmt.Println("测试完成")
}
