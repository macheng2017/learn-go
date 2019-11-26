package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	//开始将种子(request)放入到队列当中
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 将传递进来的多个种子(如果有),从队列中取出,使用相应的解析器解析
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error searching url %s; %v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)

		// 将获取到的requests重新放到队列中
		// 获取到的结果items打印出来
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
