package main //包名
import (
	"myScoket"
)

func main() {

	go func() {
		myScoket.RunTCP()
	}()
	go func() {
		myScoket.RunUDP()
	}()

	ch := make(chan int)
	ch <- 1
}
