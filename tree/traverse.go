package tree

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