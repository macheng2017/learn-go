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
		select {
		// 这样有缺点,在select中收一个数,收完之后下面一行又会阻塞
		case n := <-c1:
			w <- n
		case n := <-c2:
			w <- n

		}
	}
}
