package main //包名

import (
	"fmt"
	"my_go/algorithm/my_stack"
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
	test := my_stack.LinkStack{}
	test.Test()
	fmt.Println("OK ")

}
