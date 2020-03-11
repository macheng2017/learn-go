package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	var c1, c2 = generator1(), generator1()
	//n1 := <-c1
	//n2 := <-c2
	// 现在的问题是,我想让这两个channel谁先到谁先收数据,怎么解决?

	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
			//default:
			//	fmt.Println("No value received")
		}

	}

}
