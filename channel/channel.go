package main

import "fmt"

func channelDemo() {
	// var c chan int // c == nil 这样定义也可以但是,没有办法直接使用,
	// 问题是没法用这种定义方式有什么意义?
	// 在使用select的时候有用.
	// 使用 make 可以初始化一个channel
	c := make(chan int)
	c <- 1
	c <- 2
	n := <-c
	fmt.Println(n)
}

func main() {
	// 直接使用的结果是deadlock
	channelDemo()
	//	fatal error: all goroutines are asleep - deadlock!
}
