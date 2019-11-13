package main

import (
	"learngo/filelistingserver/filelisting"
	"net/http"
)

func main() {
	http.HandleFunc("/list/", filelisting.HandleFileList)
	serve := http.ListenAndServe(":8888", nil)

	if serve != nil {
		panic(serve)
	}
}
