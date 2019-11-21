package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	request, e := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)

	if e != nil {
		fmt.Println(e)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	// 自定义一个client,用于查看重定向信息
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			fmt.Printf("%s\n", via)
			return nil
		},
	}
	// 将DefaultClient替换成自己定义的client
	resp, err := client.Do(request)
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
