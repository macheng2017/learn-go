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
	var c1, c2 = work1(), work1()

	select {
	case n := <-c1:
		fmt.Println(n)
	case n := <-c2:
		fmt.Println(n)
		//default:
		//	fmt.Println("no value received")
	}

}
