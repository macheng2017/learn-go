package learngo

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	//s1 := arr[:6]
	//s2 := arr[2:]
	//s3 := arr[:]

	//fmt.Println("s", s)
	//fmt.Println("s1", s1)
	//updateSlice(s1)
	//fmt.Println("s1", s1)
	//fmt.Println("arr",arr)
	//fmt.Println("s2", s2)
	//updateSlice(s2)
	//fmt.Println("s2", s2)
	//fmt.Println("arr",arr)
	//fmt.Println("Reslice")
	//fmt.Println("s3", s3)
	//s3 = s3[:5]
	//fmt.Println("s3", s3)
	//s3= s3[2:]
	//fmt.Println("s3", s3)

	//fmt.Println("Extending slice")
	//fmt.Println("arr=", arr)
	//s1 = arr[2:6]
	//s2 = s1[3:5]
	//fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1, len(s1), cap(s1))
	//fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",s2, len(s2), cap(s2))

	/**
	  Extending slice
	  arr= [0 1 2 3 4 5 6 7]
	  s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	  s2=[5 6], len(s2)=2, cap(s2)=3
	*/

	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s2", s2)
	fmt.Println("arr", arr)
	s3 := append(s2, 10)
	fmt.Println("s3", s3)
	s4 := append(s3, 11)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Println("s4", s4)
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))
	s5 := append(s4, 12)
	fmt.Println("s5", s5)
	s6 := append(s5, 13)
	fmt.Printf("s6=%v, len(s6)=%d, cap(s6)=%d\n", s6, len(s6), cap(s6))
	s7 := append(s6, 13)
	fmt.Printf("s7=%v, len(s7)=%d, cap(s7)=%d\n", s7, len(s7), cap(s7))

	/**
	  s2 [5 6]
	  arr [0 1 2 3 4 5 6 7]
	  s3 [5 6 10]
	  s3=[5 6 10], len(s3)=3, cap(s3)=3
	  s4 [5 6 10 11]
	  s4=[5 6 10 11], len(s4)=4, cap(s4)=6
	  s5 [5 6 10 11 12]
	  s6=[5 6 10 11 12 13], len(s6)=6, cap(s6)=6
	  s7=[5 6 10 11 12 13 13], len(s7)=7, cap(s7)=12
	*/

	arr2 := [...]int{4, 5, 6, 7, 8}
	//arrays3(&arr2)
	// 在函数中传递一个数组参数,可以使用slice代替指针
	arrays3(arr2[:])

	fmt.Println(arr2)
}

//需要注意看的地方:
// func updateSlice(s [] int) {
// 参数列表中 s [] int,这里定义了一个slice 切片
// go语言所有类型都是值类型,但是slice 不是一个值类型,slice内部有一个数据结构,文档说slice是对arr一个view
// 上面的例子说明了一个问题:
// 1. slice 确实不是值类型,而是一个引用,因为slice修改了初始arr,而不是arr的copy
// 2. 但是go语言的参数传递不是只有值传递吗? 还是上面的参数列表中的玄机,参数列表中定义了一个切片,内部实现是对arr的一个view,
// 这个view可以修改原始arr

//func arrays3(arr *[5]int) {
func arrays3(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Printf("index: %d value: %d \n", i, v)
	}
}
