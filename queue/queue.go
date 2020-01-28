package queue

// Queue原来的类型是 []int 现在改为interface{}
type Queue []interface{}

// 方法开头大写，因为需要给外面用
// 参数使用指针类型的，这样就直接改变值，而不是拷贝一份再改变
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// push和pop虽然存的是interface{} slice 但是我们希望
// 它只接收 int 类型
func (q *Queue) Push2(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop2() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) Push3(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Queue) Pop3() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
