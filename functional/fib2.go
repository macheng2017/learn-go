package main

import "fmt"

func fibonacci1() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci1()
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	//fmt.Println(fibonacci1()())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// 这里有个问题就是,有两种调用方式 第一种就是,使用fibonacci1()()
	//第二种就是保存到一个变量当中然后使用这个变量调用返回的函数
	// f: = fibonacci1()
	// fmt.Println(f())
	// 这两种有什么区别?
	// 第二种调用 由于引用被存放到了变量当中, a,b 的值同时(注意同时,相当于一个快照,go语言的值传递)被读取到了函数当中,
	// 如果是其他语言比如java中,先读取a的值,a=b 修改方法外部的a=1 然后再,a+b 这时候方法外部的已经被修改为了a=1 ,这里a=a+b a=2
}
