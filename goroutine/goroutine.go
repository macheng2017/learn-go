package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// 这里不传值的话就会出现闭包,然后下标越界
		go func() {
			for {
				a[i]++
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
	//panic: runtime error: index out of range [10] with length 10
}
