package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

var HandleFileList = func(writer http.ResponseWriter, request *http.Request) error {
	// 给用户提示输入的url有问题
	//fmt.Println("request.URL.Path", request.URL.Path)
	a := strings.Index(request.URL.Path, prefix)
	if a != 0 {
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	fmt.Println(path)
	file, e := os.Open(path)
	if e != nil {

		return e
	}
	//fmt.Println(file)
	defer file.Close()
	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		return e
	}
	writer.Write(bytes)
	return nil
}
