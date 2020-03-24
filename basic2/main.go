package main

import "fmt"

func main() {
	var i = "nihao"

	fmt.Println([]byte(i))
	fmt.Printf("%v", []byte(i))
}

// 需要注意的是: byte 切片打印出来的是和nihao的每个字母一一对应
// [110 105 104 97 111]
