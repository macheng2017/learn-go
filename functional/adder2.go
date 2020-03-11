package main

import "fmt"

func add() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	a := add()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... +%d = %d\n", i, a(i))
	}
}
