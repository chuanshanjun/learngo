package main

import "fmt"

// go语言是没有构造函数的
// 如果需要的话可以使用工厂函数
type treeNode struct {
	value       int
	left, right *treeNode
}

// treeNode的工厂函数
func createNode(value int) *treeNode {
	// 返回回去的算是局部变量
	return &treeNode{value: value}
}

// go语言中给结构体定义方法，是在方法名前面加个接受不了者
// 就像其他语言的this
// 该函数与其他函数没啥区别
// 强调还是值传递，go语言中都是值传递
func (node treeNode) print() {
	fmt.Print(node.value)
}

func (node treeNode) setValue(value int) {
	node.value = value
}

// go语言加不加星号都是点
func (node *treeNode) setValueByPtr(value int) {
	node.value = value
}

func (node *treeNode) setValueByPtrAndCheck(value int) {
	if node == nil {
		fmt.Println("Setting value to nil" + "node. Ignored")
		// nil的话是没有值的拿不到value会报错
		// 所以这边直接return
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	// 此处做个比较
	// 如果是java或者c++ 就要再判断下，但go语言不需要
	//if node.left != null {
	//
	//}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	// 定义一个treeNode类型
	var root treeNode
	// 初始值value是0 left,right是nil
	fmt.Println(root)

	// go语言是没有构造函数的，下面这些方法就可以直接用
	// 建立树的节点
	// 给root的value赋值3
	root = treeNode{value: 3}
	// 给root装上一个左子树，且左子树都是zero value
	root.left = &treeNode{}
	// 给root装上一个右子树，且右子树赋予初始值，右子树是指针类型，所以传地址
	root.right = &treeNode{5, nil, nil}
	// 使用内建函数new() 新建一个右子树的左子树
	// 在go语言中，不管是指针还是实例，都是.下去的
	root.right.left = new(treeNode)

	// 当用slice批量赋值的时候，slice中的treeNode都可以省略
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	fmt.Println("===>>>")

	root.left.right = createNode(2)

	fmt.Println(root)

	// 调用node的print()函数
	root.print()
	fmt.Println()

	// setValue的地址改不掉,因为方法中的参数是值传递
	root.setValue(200)
	fmt.Println(root)

	// 方法参数接受一个指针即可，修改原来的值
	// 这边调用不用管，参数是不是指针
	// 实际调用的时候会把 root.right.left的地址传到方法中
	// 指针调用的话，直接把地址传给它
	root.right.left.setValueByPtr(300)
	fmt.Println(root)

	// root.right.left本来也是一个地址，当我们调用print()的时候
	// 会把这个地址解析出来，访问这个地址，把里面真正的内容拿出来，去拷一份给方法中
	root.right.left.print()

	pRoot := &root
	pRoot.print()
	pRoot.setValueByPtr(200)
	pRoot.print()

	var pRoot2 *treeNode
	pRoot2.setValueByPtrAndCheck(200)
	fmt.Println(pRoot2)
	pRoot2 = &root
	pRoot2.setValueByPtrAndCheck(300)
	pRoot2.print()

	fmt.Println("\n===>>>")

	// root是一个值，一样可以调用指针的方法
	root.traverse()
}
