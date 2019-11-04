package queue

type Queue []int

// https://golang.org/pkg/builtin/#append
// 关于append为什么要用自身的变量去接返回值这里有详细的介绍
func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)

}

func (q *Queue) Dequeue() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
