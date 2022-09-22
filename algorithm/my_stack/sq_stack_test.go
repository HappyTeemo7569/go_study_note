package my_stack

import (
	"fmt"
	"testing"
)

func Test_SqStack(t *testing.T) {

	fmt.Println("测试开始")

	my_list := new(SqStack)
	my_list.InitStack()

	for i := 1; i <= 10; i++ {
		my_list.Push(ElemType(i*i + 1))
		my_list.Echo()
	}

	var e ElemType

	my_list.Pop(&e)
	fmt.Println("删除头元素:", e)
	my_list.Echo()

	my_list.GetTop(&e)
	fmt.Println("获取栈顶:", e)

	fmt.Println("开始清空")
	my_list.ClearStack()
	if my_list.StackEmpty() {
		fmt.Println("已清空")
		my_list.Echo()
	}

	fmt.Println("测试完成")
}
