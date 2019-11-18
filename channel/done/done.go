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
		done <- true
	}

	// 为什么这句话放到循环之外就报错了?

}

func channelDemo() {

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// 向channel发数据
	for i, worker := range workers {
		worker.in <- 'a' + i

	}

	for i, worker := range workers {
		worker.in <- 'A' + i

	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	channelDemo()
}

//receive value via channel id : 9 value j
//receive value via channel id : 1 value b
//receive value via channel id : 3 value d
//receive value via channel id : 6 value g
//receive value via channel id : 5 value f
//receive value via channel id : 8 value i
//receive value via channel id : 0 value a
//receive value via channel id : 7 value h
//receive value via channel id : 2 value c
//receive value via channel id : 4 value e
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [chan send]:
//main.channelDemo()
//	/Users/mac/github/go/src/learngo/channel/done/done.go:50 +0x16a
//main.main()
//	/Users/mac/github/go/src/learngo/channel/done/done.go:62 +0x20
