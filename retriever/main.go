package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	// 当前程序不知道Retriever是什么东西,单词故意拼错的
	fmt.Println(download(mock.Retriever{"this is fack google.com"}))
	fmt.Println(download(real2.Retriever{}))
}
