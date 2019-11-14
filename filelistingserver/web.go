package main

import (
	"learngo/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

// 定义一个名字为 appHandler 的函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 为HandleFileList 包装一层错误处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// 可以使用上一节的recover来接住这个报错信息,并处理
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			// 添加控制台错误打印提示
			log.Printf("Error handling request: %s", err.Error())

			// 在这里使用刚刚定义的用户级错误信息
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

// 定义一个给用户看的错误接口
type userError interface {
	error
	Message() string
}

func main() {
	// 这个时候用户再访问 localhost:8888 就会报错
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	serve := http.ListenAndServe(":8888", nil)

	if serve != nil {
		panic(serve)
	}
}
