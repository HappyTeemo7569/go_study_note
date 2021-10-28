package main

import (
	"fmt"
	"time"
)

func startTimer(f func(now time.Time)) chan bool {
	done := make(chan bool, 1)
	go func() {
		t := time.NewTimer(time.Second * 3)
		defer t.Stop()
		select {
		case now := <-t.C:
			f(now)
		case <-done:
			fmt.Println("done")
			return
		}
	}()
	return done
}

func main() {
	done := startTimer(func(now time.Time) {
		fmt.Println(now)
	})

	//提前终止
	done <- false

	time.Sleep(5 * time.Second)
	close(done)
}
