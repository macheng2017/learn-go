package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

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

	// 2. 自动判断网页内容编码格式
	// 添加 go get  -v golang.org/x/net/html
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	bytes, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", bytes)
	printCityList(bytes)
}

func printCityList(bytes []byte) {
	// 这里的[^>]*  ^>意思是取反,不包括'>' 合起来意思是不包括'>'的任意多个字符
	// 用来匹配<a href="http://www.zhenai.com/zhenghun/changde" data-v-5e16505f>常德</a>
	// 正确的写法就是,先把上面要匹配的内容复制到字符串模板当中,把其中不同的部分改成正则表达式即可
	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+" [^>]*>[^>]+</a>`)
	all := re.FindAll(bytes, -1)
	for _, m := range all {
		fmt.Printf("%s\n", m)
	}
	fmt.Println("count tage", len(all))

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
