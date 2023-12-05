package tree

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

// 为treenode提供Print函数
func (node Node) Print() {
	fmt.Printf("%d ", node.Val)
}

// 传入指针来进行引用传递
func (node *Node) SetVal(val int) {
	if node == nil {
		fmt.Println("Setting val to nil " +
			"node . Ignored.")
		return
	}
	node.Val = val
}

// 工厂函数
func CreateNode(val int) *Node {
	// 返回局部变量的地址, 是合法的
	return &Node{Val: val}
}

// func main() {
// 	root := treeNode{val: 3}
// 	root.left = &treeNode{}
// 	root.right = &treeNode{5, nil, nil}
// 	root.right.left = new(treeNode)
// 	root.left.right = createNode(2)
// 	root.right.left.setVal(4)

// 	// root.print()
// 	// fmt.Println()

// 	// root.setVal(4)
// 	// root.print()
// 	// fmt.Println()

// 	// var pRoot *treeNode // nil指针
// 	// pRoot.setVal(200)
// 	// // pRoot.print()	// error
// 	// pRoot = &root
// 	// pRoot.setVal(300)
// 	// pRoot.print()

// 	root.traverse()
// 	// nodes := []treeNode{
// 	// 	{val: 3},
// 	// 	{},
// 	// 	{6, nil, nil},
// 	// }
// 	// fmt.Println(nodes)

// }
