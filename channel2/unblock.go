package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work1() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			i++
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			// 这样当select 收到通信的时候就终止了goroutine
			if i == 20 {
				c <- i
			}
			fmt.Println(i)
		}
	}()
	return c
}

func main() {
	var c1 = make(chan interface{})
	close(c1)
	var c2 = make(chan interface{})
	close(c2)
	var c3 = make(chan interface{})
	close(c3)
	var c1Count, c2Count, c3Count int
	for i := 10000000; i >= 0; i-- {

		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		case <-c3:
			c3Count++
			//default:
			//	fmt.Println("no value received")
		}

	}
	fmt.Printf("c1=%v, c2=%v c3=%v", c1Count, c2Count, c3Count)
}
