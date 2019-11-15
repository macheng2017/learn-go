package main

import (
	"fmt"
	"time"
)

// goroutine 会随着主线程结束而结束
func main() {
	//fmt.Println("go")
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	//fmt.Println("end")
	time.Sleep(time.Millisecond)
	fmt.Println(a)

}
