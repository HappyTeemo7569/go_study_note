package main

import "fmt"

func test01() int { //方式1
	return 250
}

//官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差
func test02() (value int) { //方式2, 给返回值命名
	value = 250
	return value
}

func test03() (value int) { //方式3, 给返回值命名
	value = 250
	return
}

//多个返回值
//求2个数的最小值和最大值
func MinAndMax(num1 int, num2 int) (min int, max int) {
	if num1 > num2 { //如果num1 大于 num2
		min = num2
		max = num1
	} else {
		max = num2
		min = num1
	}

	return
}

func main() {
	v1 := test01() //函数调用
	v2 := test02() //函数调用
	v3 := test03() //函数调用
	fmt.Printf("v1 = %d, v2 = %d, v3 = %d\n", v1, v2, v3)

	min, max := MinAndMax(33, 22)
	fmt.Printf("min = %d, max = %d\n", min, max) //min = 22, max = 33
}
