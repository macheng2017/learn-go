package main

import "fmt"

// defer 是栈结构先入后出
// defer 会在return panic之前 回到defer定义位置执行完剩余的代码
func tryDefer() {
	defer fmt.Println(2)
	defer fmt.Println(1)
	//defer deferTest()
	fmt.Println(3)
	panic("error")
	fmt.Println(4)

}

func deferTest() {
	fmt.Println("deferTest")
}

func main() {
	tryDefer()
}

// 发现一个奇怪的现象,使用panic之后 错误打印中间执行了3,1,2,而且还是随机的这些数字顺序不变但插入错误信息的位置会变动
//anic: error
//3
//
//goroutine 1 [running]:
//1
//main.tryDefer()
//2
///Users/mac/github/go/src/learngo/defer/defer.go:12 +0x1ab
//main.main()
///Users/mac/github/go/src/learngo/defer/defer.go:22 +0x20
