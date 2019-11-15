package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"aaaabbbb", 2},
		{"abcabc", 3},

		// Edge cases
		{"", 0},
		{"nil", 3},
		{"b", 1},
		{"bbbbbbbbbbbbb", 1},
		// 这个测试没问题,是我想错了,程序的目的是寻找最长不重复字符,这个字符可以从中间部分开始
		{"abcabcabcd", 4},
		// chinese support
		{"是否递四方速递", 6},
		{"师傅的说法方法", 6},
		{"aaaa的第三方", 5},
	}

	for _, tt := range tests {
		if actual := cases(tt.s); actual != tt.ans {
			t.Errorf("expected %d got %d for input %s", tt.ans, actual, tt.s)
		}
	}
}

// 性能测试
// 要写成BenchmarkXxx
func BenchmarkSubstr(b *testing.B) {
	s := "师傅的说法方法"
	// 加长字符串
	for i := 0; i < 13; i++ {
		s += s
	}
	ans := 6
	for i := 0; i < b.N; i++ {

		if actual := cases(s); actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}

//goos: darwin
//goarch: amd64
//pkg: learngo/nonDuplicate
//BenchmarkSubstr-12    	  464334	      2441 ns/op
//PASS
