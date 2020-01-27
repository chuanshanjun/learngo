package main

import (
	"fmt"

	"github.com/chuanshan/learngo/queue"
	"golang.org/x/tools/container/intsets"
)

func testSpare() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}

func main() {
	// Queue{}中给的值就是给[]int 里面的值
	// 这个q和最后的q不是一个slice
	// q.pop或者 q.push本身会改变这个q的值
	// 因为这个函数传递的是指针，指针接收者是可以改变里面的值的
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	testSpare()
}
