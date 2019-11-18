package main

import "fmt"

func main() {
	var c1, c2 chan int
	// 使用select+ default ,做一个非阻塞式收数据
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
		default:
			fmt.Println("no value received  ")
		}
	}

}
