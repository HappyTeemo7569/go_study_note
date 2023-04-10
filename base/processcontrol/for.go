package main

import "fmt"

func main() {

	//for 初始化条件 ; 判断条件 ; 条件变化

	var i, sum int

	//1) 初始化条件  i := 1
	//2) 判断条件是否为真， i <= 100， 如果为真，执行循环体，如果为假，跳出循环
	//3) 条件变化 i++
	//4) 重复2， 3， 4

	for i = 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

	Range()
}

func Range() {
	s := "abc"
	for i := range s { //支持 string/array/slice/map。
		fmt.Printf("%c\n", s[i])
	}

	for _, c := range s { // 忽略 index
		fmt.Printf("%c\n", c)
	}

	for i, c := range s {
		fmt.Printf("%d, %c\n", i, c)
	}

	for i := range s { //第2个返回值，默认丢弃，返回元素的位置(下标)
		fmt.Printf("str[%d]=%c\n", i, s[i])
	}

	for i, _ := range s { //第2个返回值，默认丢弃，返回元素的位置(下标)
		fmt.Printf("str[%d]=%c\n", i, s[i])
	}
}

func Break() {
	for i := 0; i < 5; i++ {
		if 2 == i {
			//break    //break操作是跳出当前循环
			continue //continue是跳过本次循环
		}
		fmt.Println(i)
	}

}

// Go 跳出 for-switch 和 for-select 代码块
/**
没有指定标签的 break 只会跳出 switch/select 语句，
若不能使用 return 语句跳出的话，可为 break 跳出标签指定的代码块：
*/
func Loop() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			//break    // 死循环，一直打印 breaking out...
			break loop
		}
	}
	fmt.Println("out...")
	//goto 虽然也能跳转到指定位置，但依旧会再次进入 for-switch，死循环。
}

func Goto() {

	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			goto LABEL //跳转到标签LABEL，从标签处，执行代码
		}
	}

	fmt.Println("this is test")

LABEL:
	fmt.Println("it is over")

}

// 注意：goto语句和标签之间不能有变量定义。
func gotoErr() {

	fmt.Println("start")
	goto flag
	//var say = "hello oldboy" //编译错误：goto flag jumps over declaration of say at
	//fmt.Println(say)

flag:
	fmt.Println("end")

}
