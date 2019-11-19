package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// 随机休息时间
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

//消费者
var worker = func(id int, c chan int) {
	for n := range c {
		// 避免接收空channel
		fmt.Printf("receive value via channel id : %d value %d\n", id, n)
	}
}

var createWorker = func(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	// 把收过来的n送给work
	w := createWorker(0)

	for {
		n := 0
		select {
		// 这样还有问题,就是 47,48行这两个case 收数据会等待一段时间,
		// 程序会直接执行49行,这时n恒等于0,消费者会一直打印0
		case n = <-c1:
		case n = <-c2:
		case w <- n:

		}
	}
}

//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 0
