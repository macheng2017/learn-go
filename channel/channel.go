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

func main() {
	channelDemo()
	//receive value via channel id : 8 value i
	//receive value via channel id : 5 value f
	//receive value via channel id : 6 value g
	//receive value via channel id : 1 value b
	//receive value via channel id : 0 value a
	//receive value via channel id : 3 value d
	//receive value via channel id : 2 value c
	//receive value via channel id : 7 value h
	//receive value via channel id : 9 value j
	//receive value via channel id : 4 value e
	//receive value via channel id : 4 value E
	//receive value via channel id : 2 value C
	//receive value via channel id : 0 value A
	//receive value via channel id : 3 value D
	//receive value via channel id : 6 value G
	//receive value via channel id : 7 value H
	//receive value via channel id : 5 value F
	//receive value via channel id : 8 value I
	//receive value via channel id : 9 value J
	//receive value via channel id : 1 value B
}
