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
	n := 0
	// 定义一个flag
	hasValue := false
	for {
		//即当 activeWork channel 为 nil 时select是阻塞的直到其有值
		var activeWorker chan<- int
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
