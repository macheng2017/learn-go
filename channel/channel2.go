package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	// define the channel 'c' and used make to init
	c := make(chan int)
	// channel是goroutine与goroutine之间的通信,main是一个goroutine,需要发送到另外一个goroutine接收

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	go func() {
		for {
			c <- 1
		}
	}()

	// 1. 为什么把生产者放到消费者前面会导致死锁问题
	// 2. 为什么将生产者放到goroutine当中会没有结果
	// 3. 用-race测试下,测试之后,可以正常打印,猜测问题是,时间太短了两个goroutine没时间启动就随者main goroutine结束了
	// 4. 解决方式是使用sleep函数,让main函数等一会
	//c <- 1
	//c <- 2
}

func main() {
	chanDemo()
	time.Sleep(time.Second)
} //fatal error: all goroutines are asleep - deadlock!
