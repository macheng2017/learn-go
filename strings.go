package learngo

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "aAYes我爱慕课网!"
	fmt.Println(len(s))
	fmt.Printf("%s", []byte(s))

	fmt.Printf("%x\n", []byte(s))
	for _, v := range []byte(s) {
		fmt.Printf(" %x ", v)
	} // utf-8 编码 59 65 73    e6  88  91    e7  88  b1    e6  85  95    e8  af  be    e7  bd  91    21
	// uft8 每个中文占用3个字节,采用可变长编码,每个英文字母一个字节(前三个)

	fmt.Println()
	for i, ch := range s {
		fmt.Printf(" (%d %x) ", i, ch)
	}
	fmt.Println()
	// (0 59)  (1 65)  (2 73)  (3 6211)  (6 7231)  (9 6155)  (12 8bfe)  (15 7f51)  (18 21)
	// 这里的 ch是一个 rune类型 是一个int32 类型 是一个4字节的整数, 将utf8 转成了Unicode的编码,
	// 注意中文字符开始位置索引(3 6211)  (6 7231) 也就是rune类型将,
	// 问题是: rune 是如何知道一个中文字符所占的字节数的开始位置(一个中文占三个字节)?

	// 总结: 首先将string 用utf8 解码,将解码后的每个字符转Unicode,将其存放在了rune 类型中(4字节)并打印出来

	fmt.Println("Rune count: ", utf8.RuneCountInString(s)) // Rune count:  9

	bytes := []byte(s) //bytes:  [97 65 89 101 115 230 136 145 231 136 177 230 133 149 232 175 190 231 189 145 33] 21
	// 很奇怪的是   for _,s :=range [] byte(s) 的到的结果不一样,经过计算使用range得到的是16进制
	// 但为什么一个是十进制,一个是十六进制?
	// 61  41  59  65  73  e6  88  91  e7  88  b1  e6  85  95  e8  af  be  e7  bd  91  21
	fmt.Println("bytes: ", bytes, len(bytes))
	for len(bytes) > 0 {
		fmt.Println(bytes)
		ch, size := utf8.DecodeRune(bytes)
		// DecodeRune 是取出bytes slice中的第一个utf-8编码的字符
		fmt.Printf("%c %v \n", ch, size)
		// 使用这种方式,可以删除已经取出的字符,现有的字符宽度(width)恰好是需要删除的字符的宽度
		// 注意这里的用词是字符的宽度,而不是字符串的长度,字符的宽度是指,一个字符所占用字节数,例如26个字母占用一字节,而汉字则在utf-8编码下
		// 占用2~4个字节
		bytes = bytes[size:]
		//fmt.Printf("%c ", ch)
	}

	for i, ch := range s {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
