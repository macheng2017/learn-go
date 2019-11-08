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
