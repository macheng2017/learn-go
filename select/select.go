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
		// 新问题,生产数据速度太快而消耗速度太慢的时候就会发生有些数据会跳过去
		time.Sleep(time.Millisecond)
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
	var worker = createWorker(0)
	// 解决办法: 可以让收到的n先存下来排队
	var values []int
	// 返回一个chan time
	tm := time.After(10 * time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			// 在这里进行排队,排完队之后让别人消耗他
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
			// 如果过了800ms没有生成数据则打印timeout
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		// 如果有值,则送过去消耗
		case activeWorker <- activeValue:
			// 送过去之后,将第一个移除
			values = values[1:]
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}

// 这样就解决了,生产者快于消费之,中间的数据会被跳过(覆盖)的问题
//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 1
//receive value via channel id : 0 value 1
//receive value via channel id : 0 value 2
//receive value via channel id : 0 value 2
//receive value via channel id : 0 value 3
//receive value via channel id : 0 value 3
//receive value via channel id : 0 value 4
