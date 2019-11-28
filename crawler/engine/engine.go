package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

// engine的职能
// 1.从数组中取出,放入队列中
// 2.从队列中一一取出,并使用fetcher拉取并转码
// 3.使用相应的解析器解析(main函数初始化engine一并传递过来的解析器函数),并将解析出来的
// 4.将解析出来并封装好的ParseResult中的Request重新追加到队列中
func Run(seeds ...Request) {
	// 1.从数组中取出,放入队列中
	var requests []Request

	//开始将种子(request)放入到队列当中
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {

		// 2.从队列中一一取出,并使用fetcher拉取并转码
		r := requests[0]
		requests = requests[1:]
		//
		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		// 4.将解析出来并封装好的ParseResult中的Request重新追加到队列中
		// 将获取到的requests重新放到队列中
		// 获取到的结果items打印出来
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

// 将Fetcher和Parser的动作提取出来
func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error searching url %s; %v", r.Url, err)
		return ParseResult{}, err
	}

	// 3.使用相应的解析器解析(main函数初始化engine打包Request传递过来的解析器函数),并将解析出来的
	return r.ParserFunc(body), nil
}
