package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	// go语言中对自定义类型的扩展支持是很好的
	value int
	lock  sync.Mutex
}

// 自定义一个类型atomicInt 并并定义两个方法
// 1. 递增自己
// 2. 获取自己的值
func (a *atomicInt) increment() {
	// 在数据读写的时候加上锁
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
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
