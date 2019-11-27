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
		// 由于这个Request是放到队列当中的,m[2]引用的值会随着循环改变,改成值传递就行了
		// 由于闭包每个函数拿到的name都不一样
		name := string(m[2])
		// 将城市下的用户名和用户id放入结构中
		result.Items = append(result.Items, "user "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 为了把用户名从上一级传递到下一级可以使用函数闭包
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParserProfile(contents, name)
			},
		})

	}
	return result
}
