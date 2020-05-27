package main

import "fmt"

func slices() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("%v", a[:0])
}

func main() {
	var i = "nihao"

	fmt.Println([]byte(i))
	fmt.Printf("%v\n", []byte(i))
	slices()
}

// 需要注意的是: byte 切片打印出来的是和nihao的每个字母一一对应
// 问题是go语言是怎么知道每个字母的Unicode编码长度?
// [110 105 104 97 111]
