package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func fibonacci() inGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 闭包: 形成闭包的要素
// 1. 在一个函数中返回值是一个函数
// 2. 返回的函数内部使用了其父函数的一些变量

// 定义一个类型
type inGen func() int

// 为类型实现接口,方便以后调用
// 这里实现的其实是,官方定义的io.Reader的接口,也就是这个接口定义在官方那(按照io.Reader的定义标准在这里实现),实现在这儿
func (g inGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)      // 将对象转成字符串
	return strings.NewReader(s).Read(p) // 将字符串写入p 当中
}

//实现这个接口

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
func main() {
	f := fibonacci()

	printFileContents(f)
}
