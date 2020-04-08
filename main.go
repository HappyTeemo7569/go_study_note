package main //包名

import (
	"fmt"
	"my_go/algorithm/my_list"
	_ "my_go/myLib"
)

func main() {

	//go func() {
	//	myScoket.RunTCP()
	//}()
	//go func() {
	//	myScoket.RunUDP()
	//}()
	//
	//ch := make(chan int)
	//ch <- 1
	test := my_list.CLinkList{}
	test.Test2()
	fmt.Println("OK ")

}
