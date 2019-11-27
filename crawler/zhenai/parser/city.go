package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range all {
		// 将城市下的用户名和用户id放入结构中
		result.Items = append(result.Items, "user "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 为了把用户名从上一级传递到下一级可以使用函数闭包
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParserProfile(contents, string(m[2]))
			},
		})

	}
	return result
}
