package main

import (
	"fmt"
	"learngo/queue"
)

func main() {
	q := queue.Queue{1}
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
	q.Enqueue("abc")
	fmt.Println(q.Dequeue())
	//1
	//2
	//false
	//3
	//true
	//abc
}
