package main

import (
	"fmt"
)

// 定义一个struct
type worker struct {
	in   chan int  // 向内部发送数据
	done chan bool // 向外部发送数据
}

var createWorker = func(id int) worker {
	// 初始化struct包含两个channel,这个很像javabean
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWork(id, w.in, w.done)

	return w
}

var doWork = func(id int, c chan int, done chan bool) {

	for n := range c {
		fmt.Printf("receive value via channel id : %d value %c\n", id, n)
	}

	// 为什么这句话放到循环之外就报错了?
	done <- true
}

func channelDemo() {

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// 向channel发数据
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
	// 通过添加通信 解决这种让main函数等待goroutine完成
	//time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
}

//receive value via channel id : 0 value a
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [chan receive]:
//main.channelDemo()
//	/Users/mac/github/go/src/learngo/channel/done/done.go:45 +0xe0
//main.main()
//	/Users/mac/github/go/src/learngo/channel/done/done.go:57 +0x20
//
//goroutine 18 [chan receive]:
//main.glob..func2(0x0, 0xc00006e060, 0xc00006e0c0)
//	/Users/mac/github/go/src/learngo/channel/done/done.go:27 +0x10d
//created by main.glob..func1
//	/Users/mac/github/go/src/learngo/channel/done/done.go:20 +0x98
