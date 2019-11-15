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
		go func(i int) { // race condition
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

// 将i传入函数中也会造成数据冲突,原因是在一边在print(a) 一边有goroutine在写
//(base) macs-iMac:goroutine mac$ go run -race goroutine.go
//==================
//WARNING: DATA RACE
//Read at 0x00c0000ae000 by main goroutine:
//main.main()
///Users/mac/github/go/src/learngo/goroutine/goroutine.go:20 +0xfb
//
//Previous write at 0x00c0000ae000 by goroutine 7:
//main.main.func1()
///Users/mac/github/go/src/learngo/goroutine/goroutine.go:15 +0x68
//
//Goroutine 7 (running) created at:
//main.main()
///Users/mac/github/go/src/learngo/goroutine/goroutine.go:13 +0xc3
//==================
//[20190 25790 33368 15180 18272 20444 8681 6679 16642 16493]
//Found 1 data race(s)
//exit status 66
