package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var HandleFileList = func(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[len("/list/"):]
	fmt.Printf(path)
	file, e := os.Open(path)
	if e != nil {
		http.Error(writer, e.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(file)
	defer file.Close()
	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		panic(e)
	}
	writer.Write(bytes)
}
