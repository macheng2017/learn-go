package main

import "fmt"

func main() {

	var c1, c2 chan int
	//n1 := <-c1
	//n2 := <-c2
	// 现在的问题是,我想让这两个channel谁先到谁先收数据,怎么解决?

	select {
	case n := <-c1:
		fmt.Println("Received from c1:", n)
	case n := <-c2:
		fmt.Println("Received from c2:", n)
	default:
		fmt.Println("No value received")

	}
}
