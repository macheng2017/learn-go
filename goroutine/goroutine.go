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
		go func() {
			// 强制go语言只使用一个核,就会出现goroutine不交出控制权的现象
			//runtime.GOMAXPROCS(1)
			for {
				a[i]++
				// 主动交出控制权
				//runtime.Gosched()
			}
		}()
	}
	//fmt.Println("end")
	time.Sleep(time.Millisecond)
	fmt.Println(a)
	//panic: runtime error: index out of range [10] with length 10
}
