package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
	age  int8
}

var AAA = user{name: "AAA", age: 25}
var BBB = &AAA

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "B-B-B"
}
func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

//写入管道的数据已经确定了，和消息队列一样，后面再改也没用

func main() {
	c := make(chan *user, 5)

	c <- BBB
	fmt.Println(BBB) //&{AAA 25}

	// 先传入BBB 再修改 BBB
	BBB = &user{name: "BBB", age: 100}
	go modifyUser(BBB) //modifyUser Received Vaule &{BBB 100}  //这个时候BBB已经改了

	go printUser(c) //printUser goRoutine called &{AAA 25}  //但是写入管道的还是改之前的BBB

	time.Sleep(3 * time.Second)
	fmt.Println(BBB) //&{B-B-B 100}

	fmt.Println(AAA) //{AAA 25}

}
