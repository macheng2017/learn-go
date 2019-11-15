package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// 这里不传值的话就会出现闭包,然后下标越界
		// 如果把值传递进函数当中就会将i cp 到函数中避免下标越界
		go func() { // race condition
			for {
				a[i]++
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
	//panic: runtime error: index out of range [10] with length 10
}

// 使用 go run -race goroutine.go 检测数据冲突
//==================
//WARNING: DATA RACE
//Read at 0x00c0000b6008 by goroutine 7:
//main.main.func1()
///Users/mac/github/go/src/learngo/goroutine/goroutine.go:15 +0x6b
//
//Previous write at 0x00c0000b6008 by main goroutine:
//main.main()
///Users/mac/github/go/src/learngo/goroutine/goroutine.go:11 +0x11b
//
//Goroutine 7 (running) created at:
//main.main()
///Users/mac/github/go/src/learngo/goroutine/goroutine.go:13 +0xf1
//==================
