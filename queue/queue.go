package queue

type Queue []int

// 方法开头大写，因为需要给外面用
// 参数使用指针类型的，这样就直接改变值，而不是拷贝一份再改变
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
