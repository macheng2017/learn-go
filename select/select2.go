package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 消费者工厂
var createWorker1 = func(id int) chan<- int {
	c := make(chan int)
	go worker1(id, c)

	return c
}

// 消费者
var worker1 = func(id int, c chan int) {
	// 使用这种方式也可以避免空channel
	// 通过这里的用法重新审视 range的意义
	// range 可以用于遍历 slice 例如 for i := range []byte(s) {}
	for n := range c {
		// 避免接收空channel
		time.Sleep(time.Second)
		fmt.Printf("receive value via channel id : %d value %d\n", id, n)
	}
}

// 生成者
func generator1() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	// 调度器
	var c1, c2 = generator1(), generator1()
	// 利用nil channel 的性质改造select,即当 activeWork channel 为 nil 时select是阻塞的直到其有值
	worker := createWorker1(0)
	//n1 := <-c1
	//n2 := <-c2
	// 现在的问题是,我想让这两个channel谁先到谁先收数据,怎么解决?
	// 定时器的用法,定时退出
	after := time.After(time.Second * 10)
	n := 0
	var values []int
	for {
		//即当 activeWork channel 为 nil 时select是阻塞的直到其有值
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		// 还有一个问题: 如果生产者速度太快了,消费者速度太慢,有些数据就会被跳过去?
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
			// 生产者延迟时间超过800 毫秒打印timeout
		case <-time.After(800 * time.Millisecond):
			fmt.Printf("timeout\n")
		case <-after:
			fmt.Printf("bye")
			return
		}
	}

}
