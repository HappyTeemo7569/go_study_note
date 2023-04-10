package main

import "fmt"

/**
字符类型
*/

func main() {

	var ch1, ch2, ch3 byte
	ch1 = 'a'  //字符赋值
	ch2 = 97   //字符的ascii码赋值
	ch3 = '\n' //转义字符
	fmt.Printf("ch1 = %c, ch2 = %c, %c", ch1, ch2, ch3)

	//大写转小写，小写转大写, 大小写相差32，小写大
	fmt.Printf("大写：%d， 小写：%d\n", 'A', 'a') //大写：65， 小写：97
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)
}
