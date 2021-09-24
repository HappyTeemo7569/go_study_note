package main //包名

import (
	"my_go/goroutine"
	_ "my_go/myLib"
	"my_go/web"
)

func webTest() {
	web.RunServer()
}

func goroutineTest() {
	//goroutine.RunM()
	//goroutine.RunC()
	//goroutine.RunL()
	goroutine.RunCond()
}

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

	//test := my_stack.LinkStack{}
	//test.Test()
	//fmt.Println("OK ")

	goroutineTest()

}
