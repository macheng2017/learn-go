package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("hello", "world", "print hello world")
	flag.Parse()
	fmt.Printf("s %v", *s)
}

//猜测用法
//$ go run flagString/main.go flag- hello
//s world
//正确用法
//1.先编译    go build main.go
//2.运行编译后代码 ./main.exe -hello haha
//3.结果是命令后的value值,usage则是提示内容 s haha
