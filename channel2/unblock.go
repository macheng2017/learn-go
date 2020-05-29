package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generator
func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			i++
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			// 这样当select 收到通信的时候就终止了goroutine
			fmt.Println(i)
			c <- i

		}
	}()
	return c
}

func work1(w chan interface{}) {
	for n := range w {
		fmt.Printf("receive value via channel id: vale =%v \n", n)
	}
}

// consumer
func createWork1() chan interface{} {
	w := make(chan interface{})
	go work1(w)
	return w
}

func main() {
	var c1 = generator()
	var c2 = generator()
	w1 := createWork1()
	n := 0
	hasVale := false

	for {
		var activeWorker chan<- interface{}
		if hasVale {
			activeWorker = w1
		}
		select {
		case n = <-c1:
			hasVale = true
		case n = <-c2:
			hasVale = true
		case activeWorker <- n:
			hasVale = false
		}
	}
}
