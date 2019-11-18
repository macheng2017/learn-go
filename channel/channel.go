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
		// 用make给每个nil chan 初始化
		c[i] = make(chan int)
		go worker(i, c[i])
	}

	// 向channel发数据
	for i := 0; i < 10; i++ {
		c[i] <- 'a' + i
	}

	// 防止main函数提前执行完,结束掉整个程序
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	//receive value via channel id : 8 value i
	//receive value via channel id : 0 value a
	//receive value via channel id : 6 value g
	//receive value via channel id : 1 value b
	//receive value via channel id : 2 value c
	//receive value via channel id : 4 value e
	//receive value via channel id : 3 value d
	//receive value via channel id : 5 value f
	//receive value via channel id : 7 value h
	//receive value via channel id : 9 value j
}
