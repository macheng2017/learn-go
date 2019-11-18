package main

import (
	"fmt"
	"time"
)

// 告诉外部调用者如何使用内部定义的channel
// 在返回值定义上 这个chan 只能从外部发数据(外部只写)
var createWorker = func(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("receive value via channel id : %d value %c\n", id, <-c)
		}
	}()

	return c
}

func channelDemo() {
	// 相应的这里只能定义一个只能向内部发送数据的channel(只写)
	var c [10]chan<- int
	for i := 0; i < 10; i++ {
		// 用返回的值初始化,上面定义的channel数组
		c[i] = createWorker(i)
	}

	// 向channel发数据

	for i := 0; i < 10; i++ {
		c[i] <- 'a' + i
		// 只能向其内部发送数据
		//./channel.go:34:8: invalid operation: <-c[i] (receive from send-only type chan<- int)
		//n := <-c[i]
	}

	for i := 0; i < 10; i++ {
		c[i] <- 'A' + i
	}
	// 防止main函数提前执行完,结束掉整个程序
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	// 只发数据没人来收就会报错
	c := make(chan int)
	c <- 2
}

func main() {
	//channelDemo()
	bufferedChannel()
}

//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [chan send]:
//main.bufferedChannel(...)
//	/Users/mac/github/go/src/learngo/channel/channel.go:46
//main.main()
//	/Users/mac/github/go/src/learngo/channel/channel.go:51 +0x50
