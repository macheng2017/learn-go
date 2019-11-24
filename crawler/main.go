package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code")
		return
	}
	// 判断网页内容编码格式两种方式

	// 1. 手动判断网页编码格式
	// 添加 go get -v golang.org/x/text/
	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	bytes, err := ioutil.ReadAll(utf8Reader)

	//bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)

}

// 2. 自动判断网页内容编码格式
// 添加 go get  -v golang.org/x/net/html
func determineEncoding(r io.Reader) encoding.Encoding {
	charset.DetermineEncoding()
}
