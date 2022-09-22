package main

import "fmt"

/**
fmt包
*/

func main() {
	//整型
	a := 15
	fmt.Printf("a = %b\n", a) //a = 1111
	fmt.Printf("%%\n")        //只输出一个%

	//字符
	ch := 'a'
	fmt.Printf("ch = %c, %c\n", ch, 97) //a, a

	//浮点型
	f := 3.14
	fmt.Printf("f = %f, %g\n", f, f) //f = 3.140000, 3.14
	fmt.Printf("f type = %T\n", f)   //f type = float64

	//复数类型
	v := complex(3.2, 12)
	fmt.Printf("v = %f, %g\n", v, v) //v = (3.200000+12.000000i), (3.2+12i)
	fmt.Printf("v type = %T\n", v)   //v type = complex128

	//布尔类型
	fmt.Printf("%t, %t\n", true, false) //true, false

	//字符串
	str := "hello go"
	fmt.Printf("str = %s\n", str) //str = hello go

	//%T可以打印类型
	fmt.Printf("str的类型: %T \n", str)

	/**
	输入
	*/
	var vInput int
	fmt.Println("请输入一个整型：")
	fmt.Scanf("%d", &vInput)
	//fmt.Scan(&v)
	fmt.Println("vInput = ", vInput)

	/**
	类型转换
	*/
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	//bool类型不能转换为int
	//fmt.Printf("flag = %d\n", int(flag))

	//0就是假，非0就是真
	//整型也不能转换为bool
	//flag = bool(1)

	var ch1 byte
	ch1 = 'a' //字符类型本质上就是整型
	var t int
	t = int(ch1) //类型转换，把ch的值取出来后，转成int再给t赋值
	fmt.Println("t = ", t)

	//别名
	type (
		myint int    //int改名为myint
		mystr string //string改名为mystr
	)

	var my_a myint
	my_a = 1
	fmt.Println(my_a)
}
