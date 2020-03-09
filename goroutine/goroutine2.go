package main

import (
	"fmt"
	"time"
)

func main() {
	var count [10]int
	for i := 0; i < len(count); i++ {
		go func(i int) {
			for {
				count[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Print(count)
}

// go run -race goroutine2.go 使用-race参数可以发现冲突
// 即便改过来之后也会发现 data race 原因是goroutine 一边在写,一边在读,这个问题可以通过channel来解决

// 需要注意的问题;
// 1. 匿名函数立即执行函数
// 2. goroutine 是随着主线程执行完毕而停止
// 3. goroutine的传值方式闭包
// 4. goroutine 是非抢占式协程(意味着协程主动(程序员主动控制)交出控制权)
