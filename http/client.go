package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	// bytes, err := ioutil.ReadAll(resp.Body)
	// 与之前的相比较多了头部信息
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", bytes)

}
