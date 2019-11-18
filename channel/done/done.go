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
	// 使用channel通知main函数或其他函数,执行完毕

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
//receive value via channel id : 1 value b
//receive value via channel id : 2 value c
//receive value via channel id : 3 value d
//receive value via channel id : 4 value e
//receive value via channel id : 5 value f
//receive value via channel id : 6 value g
//receive value via channel id : 7 value h
//receive value via channel id : 8 value i
//receive value via channel id : 9 value j
//receive value via channel id : 0 value A
//receive value via channel id : 1 value B
//receive value via channel id : 2 value C
//receive value via channel id : 3 value D
//receive value via channel id : 4 value E
//receive value via channel id : 5 value F
//receive value via channel id : 6 value G
//receive value via channel id : 7 value H
//receive value via channel id : 8 value I
//receive value via channel id : 9 value J
