package main

import (
	"fmt"
	"learngo/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.google.com")
}

func main() {
	// 当前程序不知道Retriever是什么东西,单词故意拼错的
	var r Retriever
	r = mock.Retriever{"this is fack google.com"}
	fmt.Println(download(r))
}
