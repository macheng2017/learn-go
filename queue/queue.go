package queue

type Queue []interface{}

// https://golang.org/pkg/builtin/#append
// 关于append为什么要用自身的变量去接返回值这里有详细的介绍
func (q *Queue) Enqueue(v interface{}) {
	*q = append(*q, v)
	//*q = append(*q, v.(int)) //内部限定类型的方法
}

func (q *Queue) Dequeue() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
	//return head.(int) // 限定特定的返回类型
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
