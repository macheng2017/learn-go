package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}

// 需要注意的问题;
// 1. 匿名函数立即执行函数
// 2. goroutine 是随着主线程执行完毕而停止
// 3. goroutine的传值方式闭包
