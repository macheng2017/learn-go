package main

import (
	"learngo/filelistingserver/filelisting"
	"net/http"
	"os"
)

// 定义一个名字为 appHandler 的函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 为HandleFileList 包装一层错误处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {

			switch {
			case os.IsNotExist(err):
				http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			}
		}
	}
}
func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	serve := http.ListenAndServe(":8888", nil)

	if serve != nil {
		panic(serve)
	}
}
