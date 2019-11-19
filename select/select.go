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
		time.Sleep(5000 * time.Millisecond)
		fmt.Printf("receive value via channel id : %d value %d\n", id, n)
	}
}

//receive value via channel id : 0 value 0
//receive value via channel id : 0 value 7
//receive value via channel id : 0 value 12
//receive value via channel id : 0 value 20
//receive value via channel id : 0 value 27
//receive value via channel id : 0 value 34
var createWorker = func(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	n := 0
	hasValue := false
	for {
		var activeWorker chan<- int // 利用这样定义初始值为nil的性质
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}
}
