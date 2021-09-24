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

func main() {
	c := make(chan *user, 5)
	c <- BBB
	fmt.Println(BBB)
	// 先传入BBB 再修改 BBB
	BBB = &user{name: "BBB", age: 100}
	go printUser(c)
	go modifyUser(BBB)
	time.Sleep(5 * time.Second)
	fmt.Println(BBB)
}
