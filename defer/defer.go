package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

// defer 是栈结构先入后出
// defer 会在return panic之前 回到defer定义位置执行完剩余的代码
func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print to many")
		}
	}

}

func deferTest() {
	fmt.Println("deferTest")
}

func writeFile(filename string) {
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	tryDefer()
	//writeFile("test.txt")
}

//30
//panic: print to many
//29
//
//28
//goroutine 1 [running]:
//27
//main.tryDefer()
//26
///Users/mac/github/go/src/learngo/defer/defer.go:16 +0x111
//25
//main.main()
//24
///Users/mac/github/go/src/learngo/defer/defer.go:42 +0x20
//23
//22
//21
//20
//19
//18
//17
//16
//15
//14
//13
//12
//11
//10
//9
//8
//7
//6
//5
//4
//3
//2
//1
//0
