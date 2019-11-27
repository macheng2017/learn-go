package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	// 通过加入header来避免403错误(反爬机制)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code")
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	// 判断网页内容编码格式两种方式

	// 2. 自动判断网页内容编码格式
	// 添加 go get  -v golang.org/x/net/html
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

// 有个问题,使用了determineEncoding 之后打印内容不全,这是怎么回事?
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 这里使用peek之后将其存储起来,外面再次transform.NewReader 是从1025的位置读取的,所以前面的好像被截掉了
	bytes, err := r.Peek(1024)
	if err != nil {
		// 如果解析错误则打印错误日志,还是返回一个默认的编码格式
		log.Panicf("Fetcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
