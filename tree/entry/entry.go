package main

import (
	"fmt"
	"github.com/chuanshan/learngo/tree"
)

// 定义一个myTreeNode的结构体
// 名字用小写，因为是自己用
// 通过组合的方法来扩展Node
type myTreeNode struct {
	// 值的类型是指针，这样就不需要直接拷贝一份了
	node *tree.Node
}

// 给结构体定义方法,也用小写的，因为也是自己用
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

// 一个目录下是一个包
func main() {

	// 定义一个treeNode类型
	var root tree.Node
	// 初始值value是0 left,right是nil
	fmt.Println(root)

	// go语言是没有构造函数的，下面这些方法就可以直接用
	// 建立树的节点
	// 给root的value赋值3
	root = tree.Node{Value: 3}
	// 给root装上一个左子树，且左子树都是zero value
	root.Left = &tree.Node{}
	// 给root装上一个右子树，且右子树赋予初始值，右子树是指针类型，所以传地址
	root.Right = &tree.Node{5, nil, nil}
	// 使用内建函数new() 新建一个右子树的左子树
	// 在go语言中，不管是指针还是实例，都是.下去的
	root.Right.Left = new(tree.Node)

	// 当用slice批量赋值的时候，slice中的treeNode都可以省略
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	fmt.Println("===>>>")

	root.Left.Right = tree.CreateNode(2)

	fmt.Println(root)

	// 调用node的print()函数
	root.Print()
	fmt.Println()

	// setValue的地址改不掉,因为方法中的参数是值传递
	root.SetValue(200)
	fmt.Println(root)

	// 方法参数接受一个指针即可，修改原来的值
	// 这边调用不用管，参数是不是指针
	// 实际调用的时候会把 root.right.left的地址传到方法中
	// 指针调用的话，直接把地址传给它
	root.Right.Left.SetValueByPtr(300)
	fmt.Println(root)

	// root.right.left本来也是一个地址，当我们调用print()的时候
	// 会把这个地址解析出来，访问这个地址，把里面真正的内容拿出来，去拷一份给方法中
	root.Right.Left.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValueByPtr(200)
	pRoot.Print()

	var pRoot2 *tree.Node
	pRoot2.SetValueByPtrAndCheck(200)
	fmt.Println(pRoot2)
	pRoot2 = &root
	pRoot2.SetValueByPtrAndCheck(300)
	pRoot2.Print()

	fmt.Println("\n===>>>")

	// root是一个值，一样可以调用指针的方法
	root.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
