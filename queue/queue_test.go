package queue

import "fmt"

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push2(2)
	q.Push2(3)
	fmt.Println(q.Pop2())
	fmt.Println(q.Pop2())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop2())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}
