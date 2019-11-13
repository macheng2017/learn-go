package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		fmt.Printf(path)
		file, e := os.Open(path)
		if e != nil {
			panic(e)
		}
		//fmt.Println(file)
		defer file.Close()
		bytes, e := ioutil.ReadAll(file)
		if e != nil {
			panic(e)
		}
		writer.Write(bytes)
	})
	serve := http.ListenAndServe(":8888", nil)

	if serve != nil {
		panic(serve)
	}
}
