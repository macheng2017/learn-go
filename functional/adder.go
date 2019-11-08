package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 使用一个递归
// 递归三原则:
// 1. 递归是调用自己的函数
// 2. 递归函数中有一个出口条件
// 3. 递归函数是逐步到达出口条件的
// 由此可见上面的闭包与递归有本质的区别
func adder2(arr []int, l int) int {
	if l == len(arr) {
		return 0
	}
	return arr[l] + adder2(arr, l+1)
}

// user
func sum(arr []int) int {
	return adder2(arr, 0)
}

//  下面的这个"递归" 没有出口条件,太奇怪了
// 对一些新东西,当自己感觉很不适应,或者感觉奇怪的时候,说明已经达到了自己的认知边界上了
// ,这时候应该为如果弄懂这个就会延伸自己的认知疆域而感到高兴而不是沮丧
type iAdder func(int) (int, iAdder)

func adder3(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder3(base + v)
	}
}

func main() {
	//a := adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(a(i))
	//}
	const size = 100
	arr := [size]int{}
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	//fmt.Println(arr)
	fmt.Println(sum(arr[:]))
}
