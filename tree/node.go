package tree

import "fmt"

// go语言是没有构造函数的
// 如果需要的话可以使用工厂函数
type Node struct {
	Value       int
	Left, Right *Node
}

// treeNode的工厂函数
func CreateNode(value int) *Node {
	// 返回回去的算是局部变量
	return &Node{Value: value}
}

// go语言中给结构体定义方法，是在方法名前面加个接受不了者
// 就像其他语言的this
// 该函数与其他函数没啥区别
// 强调还是值传递，go语言中都是值传递
func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node Node) SetValue(value int) {
	node.Value = value
}

// go语言加不加星号都是点
func (node *Node) SetValueByPtr(value int) {
	node.Value = value
}

func (node *Node) SetValueByPtrAndCheck(value int) {
	if node == nil {
		fmt.Println("Setting value to nil" + "node. Ignored")
		// nil的话是没有值的拿不到value会报错
		// 所以这边直接return
		return
	}
	node.Value = value
}


