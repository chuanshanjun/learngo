package tree

import "fmt"

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	// 此处做个比较
	// 如果是java或者c++ 就要再判断下，但go语言不需要
	//if node.left != null {
	//
	//}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 我们在调用TraverseNew的时候会调用
// TraverseFunc的参数是一个函数，我们可以在这个函数中写上
// 我们需要遍历的时候做些什么操作
func (node *Node) TraverseNew() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

// 我们改造Traverse()方法,目前的树的遍历只是打印
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
