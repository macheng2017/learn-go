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
			panic(r)
		}
	}()

	// ...

	// panic 向上throw一个错误
	//panic(errors.New("this is a error "))
	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
	// Error occurred runtime error: integer divide by zero
}
