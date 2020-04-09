package main

import (
	"fmt"
	"runtime"
)

func main() {

	// Channel Buffered mirip seperti channel, proses serah terima data,
	// namun channel buffered memungkinkan channel dapat membatasi proses pengiriman data ke channel
	// note : proses pengiriman channel dilakukan dengan asynchronous, namun proses penerimaan/mengambil channel dengan synchronous

	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages // proses penerimaan channel
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i // proses pengiriman channel
	}

}
