package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func retrieve(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

func main() {
	fmt.Println(retrieve("https://www.baidu.com"))
}
