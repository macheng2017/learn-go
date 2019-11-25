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
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)

	s := re.FindAllString(text, -1)
	fmt.Println(s)
}
