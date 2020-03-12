package main

import (
	"fmt"
	"time"
)

type atomicInt int

// 自定义一个类型atomicInt 并并定义两个方法
// 1. 递增自己
// 2. 获取自己的值
func (a *atomicInt) increment() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}

// 2
//虽然是结果是正确的但是,数据读写冲突,需要想办法解决这个问题?
// 1. 可以用channel + select解决
// 2. 用传统的同步机制解决

//WARNING: DATA RACE
//Read at 0x00c0000ae008 by main goroutine:
//main.main()
///Users/mac/github/go/src/learngo/basic/atomic/atomic.go:18 +0xc5
//
//Previous write at 0x00c0000ae008 by goroutine 7:
//main.main.func1()
///Users/mac/github/go/src/learngo/basic/atomic/atomic.go:14 +0x51
//
//Goroutine 7 (finished) created at:
//main.main()
///Users/mac/github/go/src/learngo/basic/atomic/atomic.go:24 +0xaa
