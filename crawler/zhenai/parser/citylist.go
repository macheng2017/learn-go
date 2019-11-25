package parser

import (
	"fmt"
	engine "learngo/crawler"
	"regexp"
)

func ParseCityList(contents []byte) engine.ParseResult {
	// 这里的[^>]*  ^>意思是取反,不包括'>' 合起来意思是不包括'>'的任意多个字符
	// 用来匹配<a href="http://www.zhenai.com/zhenghun/changde" data-v-5e16505f>常德</a>
	// 正确的写法就是,先把上面要匹配的内容复制到字符串模板当中,把其中不同的部分改成正则表达式即可
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^>]+)</a>`)
	all := re.FindAllSubmatch(contents, -1)
	for _, m := range all {
		fmt.Printf(" city: %s URL: %s\n", m[2], m[1])
	}
	fmt.Println("count tage", len(all))
}
