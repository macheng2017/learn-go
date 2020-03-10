package main

import "fmt"

func chanDemo() {
	// define the channel 'c' and used make to init
	c := make(chan int)
	// channel是goroutine与goroutine之间的通信,main是一个goroutine,需要发送到另外一个goroutine接收
	c <- 1
	c <- 2
	//
	n := <-c
	fmt.Println(n)
}

func main() {
	chanDemo()
} //fatal error: all goroutines are asleep - deadlock!
