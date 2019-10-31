package learngo

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v  len(s)=%d cap(s)=%d \n", s, len(s), cap(s))
}

func main() {
	var s []int
	printSlice(s)
	for i := 0; i < 5; i++ {
		s = append(s, i*2+1)
		printSlice(s)
	}
	fmt.Println(s)
	// 空 slice的定义方式
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := make([]int, 8, 16)
	s3 := make([]int, 16, 32)
	printSlice(s2)
	printSlice(s3)
	copy(s2, s1)
	printSlice(s2)
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	/**
	Deleting elements from slice
	[1 2 3 5 6 7 8]  len(s)=7 cap(s)=16
	*/

}
