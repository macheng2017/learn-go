package learngo

import "fmt"

func poiner() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(*&a, &a, *pa)
}
func address() {
	a := 3
	fmt.Printf("变量的地址: %x\n", &a)
}

func poiner2() {
	a := 3
	var ip *int
	ip = &a
	fmt.Println("变量a 的内存地址", &a)
	fmt.Println(ip)

}

func pointer3() {
	a := 20     // 声明实际变量
	var ip *int // 声明int指针类型变量
	ip = &a     // 指针变量的存储地址

	fmt.Printf("a 变量的地址是: %x\n", &a) // 变量a的地址
	// 指针变量的存储地址
	fmt.Printf("ip 变量的存储的指针地址: %x\n", ip)
	// 使用指针访问值
	fmt.Printf("*ip 变量的值:%d\n", *ip)

}

// 交换两个值
func swap(a, b int) {
	b, a = a, b
}
func swap2(a, b *int) {
	*b, *a = *a, *b
}
func swap3(a, b int) (int, int) {
	return b, a
}

func main() {
	//poiner()
	//poiner2()
	//num := 10
	//var ip *int = &num
	//fmt.Println("ip的值", ip)
	//println()
	pointer3()
	a, b := 3, 4
	//swap(a,b)
	//swap2(&a, &b)
	a, b = swap3(a, b)
	fmt.Println(a, b)

}
