package main

import (
	"learngo/filelistingserver/filelisting"
	"net/http"
)

// 定义一个名字为 appHandler 的函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 为HandleFileList 包装一层错误处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

}
func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	serve := http.ListenAndServe(":8888", nil)

	if serve != nil {
		panic(serve)
	}
}
