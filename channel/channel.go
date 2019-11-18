package main

import (
	"fmt"
	"time"
)

func channelDemo() {

	c := make(chan int)

	go func(c chan int) {
		// 在goroutine中接收数据
		// 在这里形成了闭包,修改为下面的结构改为非闭包(值传递)
		for {
			n := <-c
			fmt.Printf("receive value via channel : %d\n", n)
		}
	}(c)

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
