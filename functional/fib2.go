package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci1() inGen1 {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type inGen1 func() int

// 使用这种写法的好处就是,可以一次性的获取一个数列输出流,不用再生成一个输出一个数字
func (g inGen1) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents1(file io.Reader) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci1()
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	printFileContents1(f)

	// 这里有个问题就是,有两种调用方式 第一种就是,使用fibonacci1()()
	//第二种就是保存到一个变量当中然后使用这个变量调用返回的函数
	// f: = fibonacci1()
	// fmt.Println(f())
	// 这两种有什么区别?
	// 第二种调用 由于引用被存放到了变量当中, a,b 的值同时(注意同时,相当于一个快照,go语言的值传递)被读取到了函数当中,
	// 如果是其他语言比如java中,先读取a的值,a=b 修改方法外部的a=1 然后再,a+b 这时候方法外部的已经被修改为了a=1 ,这里a=a+b a=2
}
