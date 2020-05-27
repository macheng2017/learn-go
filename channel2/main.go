package main

import (
	"fmt"
	"time"
)

func channel() {
	c := make(chan int)
	go func(c chan int) {
		for {
			fmt.Println(<-c)
		}
	}(c)
	i := 0
	go func() {

		for {
			i++
			c <- i
		}
	}()

}

func main() {
	channel()
	time.Sleep(time.Millisecond)
}
