package myScoket

func RunTCP() {
	go func() {
		TCPServerRun()
	}()

	go func() {
		TCPClient()
	}()

	ch := make(chan int)
	ch <- 1
}

func RunUDP() {
	go func() {
		UDPServerRun()
	}()

	go func() {
		UDPClient()
	}()

	ch := make(chan int)
	ch <- 1
}
