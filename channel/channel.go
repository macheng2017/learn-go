package main

import (
	"fmt"
	"time"
)

var worker = func(id int, c chan int) {
	// 在这里形成了闭包,修改为下面的结构改为非闭包(值传递)
	for {
		fmt.Printf("receive value via channel id : %d value %c\n", id, <-c)
	}
}

func channelDemo() {

	var c [10]chan int
	for i := 0; i < 10; i++ {
		// 初始化10个 goroutine 和10个 channel
		go worker(i, c[i])
	}

	// 向channel发数据
	for i := 0; i < 10; i++ {
		c[i] <- 'a' + 1
	}

	// 防止main函数提前执行完,结束掉整个程序
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	// 执行后发现有个错误,不能向一个 nil chan 发送数据
	//	goroutine 1 [chan send (nil chan)]:
	//main.channelDemo()
	//	/Users/mac/github/go/src/learngo/channel/channel.go:25 +0xaf
	//main.main()
	//	/Users/mac/github/go/src/learngo/channel/channel.go:33 +0x20
}
