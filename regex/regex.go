package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `my email is macheng2025@gmail.com
			email2 haha@gmail.com.cn
	email3 13222@gmail.com.cn@abc.com.cn
`
	// 提取email中的内容
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	s := re.FindAllStringSubmatch(text, -1)
	for _, m := range s {
		fmt.Println(m)
	}
}
