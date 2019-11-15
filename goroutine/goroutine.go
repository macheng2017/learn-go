package main

import (
	"fmt"
	"time"
)

// goroutine 会随着主线程结束而结束
func main() {
	fmt.Println("go")
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d \n", i)
			}
		}(i)
	}
	//fmt.Println("end")
	time.Sleep(time.Millisecond)

}
