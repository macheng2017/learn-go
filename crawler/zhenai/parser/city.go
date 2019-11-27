package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="http://album.zhenai.com/u/([0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range all {
		// 将城市下的用户名和用户id放入结构中
		result.Items = append(result.Items, "user "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})

	}
	return result
}
