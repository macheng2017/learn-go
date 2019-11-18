package main

import (
	"fmt"
	"time"
)

var worker = func(id int, c chan int) {
	// 在这里形成了闭包,修改为下面的结构改为非闭包(值传递)
	for {
		fmt.Printf("receive value via channel id : %d value %d\n", id, <-c)
	}
}

func channelDemo() {

	c := make(chan int)

	go worker(0, c)

	// 向channel发数据
	c <- 1
	c <- 2

	// 防止main函数提前执行完,结束掉整个程序
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	//receive value via channel id : 0 value 1
	//receive value via channel id : 0 value 2
	//
	//Process finished with exit code 0
}
