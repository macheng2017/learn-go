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
			c <- i
		}
	}()
	return c
}

func main() {
	var c1, c2 = work1(), work1()
	// 注意:这里有个for loop 现在的select 收到 work1的通信也不会停止会一直循环接收
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
			//default:
			//	fmt.Println("no value received")
		}
	}

}
