package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code")
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	// 判断网页内容编码格式两种方式

	// 2. 自动判断网页内容编码格式
	// 添加 go get  -v golang.org/x/net/html
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

// 有个问题,使用了determineEncoding 之后打印内容不全,这是怎么回事?
func determineEncoding(r io.Reader) encoding.Encoding {
	//直接使用 resp.body读完之后就没办法再读了?
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
