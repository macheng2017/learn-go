package main

import (
	"fmt"
	"time"
)

func main() {
	var count [10]int
	for i := 0; i < len(count); i++ {
		// 这里发生了闭包
		go func() {
			for {
				// 这里的 i 值是与函数外层for循环的i值保持一致,当最后一次循环时,i=10这里就会出错
				// 解决这个问题就是传入一个i让函数内部使用i的副本(值传递)
				count[i]++
				//runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Print(count)
}

// panic: runtime error: index out of range [10] with length 10
// go run -race goroutine2.go 使用-race参数可以发现冲突

// 需要注意的问题;
// 1. 匿名函数立即执行函数
// 2. goroutine 是随着主线程执行完毕而停止
// 3. goroutine的传值方式闭包
// 4. goroutine 是非抢占式协程(意味着协程主动(程序员主动控制)交出控制权)
