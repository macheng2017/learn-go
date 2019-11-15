package main

import "fmt"

func cases(str string) int {

	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	// 将传入的字符串转换成切片
	for i, ch := range []rune(str) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength

}
func main() {
	//fmt.Println(cases("asdfsdfdsf"))
	s := "你好啊的说法都是荆防颗粒的说法都是"

	for i, ch := range []rune(s) {
		//fmt.Println(i, ch)
		fmt.Printf("%c %d %d \n", ch, ch, i)
	}

}
