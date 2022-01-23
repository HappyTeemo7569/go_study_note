package main

import "fmt"

func main()  {
	var m1 map[int]string  //只是声明一个map，没有初始化, 此为空(nil)map
	fmt.Println(m1 == nil) //true
	//m1[1] = "mike" //err, panic: assignment to entry in nil map

	//m2, m3的创建方法是等价的
	m2 := map[int]string{}
	m3 := make(map[int]string)
	fmt.Println(m2, m3) //map[] map[]

	m4 := make(map[int]string, 10) //第2个参数指定容量
	fmt.Println(m4)                //map[]
}

