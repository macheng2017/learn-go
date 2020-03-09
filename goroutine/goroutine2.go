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
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Print(count)
}

// 需要注意的问题;
// 1. 匿名函数立即执行函数
// 2. goroutine 是随着主线程执行完毕而停止
// 3. goroutine的传值方式闭包
// 4. goroutine 是非抢占式协程(意味着协程主动(程序员主动控制)交出控制权)
