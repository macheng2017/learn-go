package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

// 跑了下测试,发现idea把保存成文件的html代码格式化了,在标签中加入了换行符,可以用\s来匹配
// 对照测试结果,看了html文件才明白原因,经验:先别着急看代码,先把简单的容易验证的环节先验证一遍再,修改代码
// 代打比方就是,电脑突然黑屏了,最容易验证的环节应该是看看显示器电源是否松了
// 空格就是' '，\s    匹配任何空白字符，包括空格、制表符、换页符等等。等价于 [ \f\n\r\t\v]。
/// /.test(' '); //true
const cityListRe = `<a\s+href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^>]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range all {
		// 将城市名和URL放入定义的结构当中
		result.Items = append(result.Items, "city "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

	}
	return result
}
