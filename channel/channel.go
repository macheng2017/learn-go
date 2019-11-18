package main

import (
	"fmt"
	"time"
)

// 告诉外部调用者如何使用内部定义的channel
// 在返回值定义上 这个chan 只能从外部发数据(外部只写)
var createWorker = func(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c
}

var worker = func(id int, c chan int) {
	// 使用这种方式也可以避免空channel
	// 通过这里的用法重新审视 range的意义
	// range 可以用于遍历 slice 例如 for i := range []byte(s) {}
	for n := range c {
		// 避免接收空channel
		fmt.Printf("receive value via channel id : %d value %c\n", id, n)
	}
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
	// 添加一个参数,缓冲区大小
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}
func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	// close 之后将channel接收到的都是空串,再接收方 使用for range 或者 判断下break即可
	time.Sleep(time.Millisecond)
}
func main() {
	//channelDemo()
	//bufferedChannel()
	channelClose()
}
