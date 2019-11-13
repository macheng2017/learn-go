package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var HandleFileList = func(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	fmt.Printf(path)
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
