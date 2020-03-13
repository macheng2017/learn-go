package main

import "fmt"

var c chan bool

func main() {

	go func() {
		for {
			fmt.Println("123")
		}
	}()
	// 使用这种方式可以阻塞当前的goroutine,从而可以让上一个goroutine一直挂起
	<-c
	//fmt.Println(<-c)

}
