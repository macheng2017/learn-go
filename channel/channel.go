package main

import (
	"fmt"
	"time"
)

func channelDemo() {

	c := make(chan int)

	// 如果使用channel发数据没人接收就会deadlock,需要用goroutine接收数据
	go func() {
		// 在goroutine中接收数据
		for {
			n := <-c
			fmt.Printf("receive value via channel : %d\n", n)
		}
	}()

	// 向channel发数据
	c <- 1
	c <- 2

	// 防止main函数提前执行完,结束掉整个程序
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	//receive value via channel : 1
	//receive value via channel : 2
	//
	//Process finished with exit code 0
}
