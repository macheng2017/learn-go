package learngo

import "fmt"

func arrays() {
	var arr [5]int
	arr1 := [3]int{2, 3, 4}
	arr2 := [...]int{4, 5, 6, 7, 8}
	var grid [4][5]int

	fmt.Println(arr, arr1, arr2, grid)
}

func arrays1() {
	arr2 := [...]int{4, 5, 6}
	for i, v := range arr2 {
		fmt.Printf("index: %d value: %d \n", i, v)
	}
}

func arrays2(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Printf("index: %d value: %d \n", i, v)
	}
}

func main() {
	//arrays()
	//arrays1()
	arr2 := [...]int{4, 5, 6, 7, 8}
	arrays2(&arr2)
	fmt.Println(arr2)
	/**
	index: 0 value: 100
	index: 1 value: 5
	index: 2 value: 6
	index: 3 value: 7
	index: 4 value: 8
	[100 5 6 7 8]

	*/
}
