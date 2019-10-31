package learngo

import (
	"fmt"
	"math"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}
func consts2() {
	const (
		a = 3
		b = 4
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("%.1f", c)
}

func enums() {
	const (
		java       = 0
		python     = 1
		javascript = 2
		golang     = 3
	)
	fmt.Println(java, python, javascript, golang)
} //0 1 2 3
func enums2() {
	const (
		java = iota
		_
		javascript
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(java, javascript, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
} // 0 2 3
//1 1024 1048576 1073741824 1099511627776 1125899906842624

func main() {
	//fmt.Printf("hello, world\n")
	//fmt.Println(aa, ss, bb)
	//triangle()
	//consts()
	//consts2()
	//enums()
	enums2()

}
