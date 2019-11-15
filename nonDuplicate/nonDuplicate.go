package main

import "fmt"

var lastOccurred = make([]int, 0xffff) // 65535

func cases(str string) int {

	//lastOccurred := make(map[rune]int)
	// 优化的原则用空间换时间
	// 对字符型的优化方法,开一个比较大的slice

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0
	// 将传入的字符串转换成切片
	for i, ch := range []rune(str) {

		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
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
