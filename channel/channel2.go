package main

import "fmt"

func chanDemo() {
	// define the channel 'c' and used make to init
	c := make(chan int)
	// channel是goroutine与goroutine之间的通信,main是一个goroutine,需要发送到另外一个goroutine接收

	//for i := 0; i < 100; i++ {
	//	c <- i
	//}
	// 为什么把生产者放到消费者前面会导致死锁问题
	c <- 1
	c <- 2
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

}

func main() {
	chanDemo()
} //fatal error: all goroutines are asleep - deadlock!
