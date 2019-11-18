package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func main() {
	var c1, c2 = generator(), generator()
	// 使用select+ default ,做一个非阻塞式收数据

	for {
		select {
		case n := <-c1:
			fmt.Println("Receive from c1", n)
		case n := <-c2:
			fmt.Println("Receive from c2", n)
			//default:
			//	fmt.Println("no value received  ")
			//
		}
	}
	//Receive from c1 0
	//Receive from c2 0
	//Receive from c2 1
	//Receive from c1 1
	//Receive from c1 2
	//Receive from c2 2
	//Receive from c1 3
	//Receive from c2 3
	//Receive from c2 4
	//Receive from c2 5
	//Receive from c1 4
	//Receive from c1 5
	//Receive from c1 6
	//Receive from c2 6
}
