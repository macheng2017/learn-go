package main

import (
	"fmt"
)

func tryRecover() {
	// 使用recover接收panic抛出的错误,recover只能在defer中定义
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(fmt.Sprintf("I dont know what is something %v", r))
		}
	}()

	// ...

	// panic 向上throw一个错误
	//panic(errors.New("this is a error "))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	// 抛出一个不是error的错误
	panic(123)
}

func main() {
	tryRecover()
	//panic: 123 [recovered]
	//panic: I dont know what is something 123
}
