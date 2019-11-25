package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "my email is macheng2025@gmail.com"
	re := regexp.MustCompile(`.+@.+\..+`)

	s := re.FindAllString(text, -1)
	fmt.Println(s)
}
