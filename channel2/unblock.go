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
		time.Sleep(time.Second)
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
	var values []int
	after := time.After(time.Second * 10)
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- interface{}
		var activeValue interface{}
		if len(values) > 0 {
			activeWorker = w1
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case <-time.After(time.Millisecond * 800):
			fmt.Println("timeout")

		case activeWorker <- activeValue:
			values = values[1:]

		case <-tick:
			fmt.Printf("Quene =%v %v\n", values, len(values))
		case <-after:
			fmt.Println("bye")
			return
		}
	}
}
