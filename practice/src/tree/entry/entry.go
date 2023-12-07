package main

import (
	"fmt"
	"learngo/src/tree"
)

// 扩展tree包中的node
// // 1. 组合
// type MyTreeNode struct {
// 	node *tree.Node
// }

// func (myNode *MyTreeNode) preOrder() {
// 	if myNode == nil || myNode.node == nil {
// 		return
// 	}

// 	left := MyTreeNode{myNode.node.Left}
// 	right := MyTreeNode{myNode.node.Right}

// 	left.preOrder()
// 	right.preOrder()
// 	myNode.node.Print()
// }

// 3. 内嵌Embedding(类似继承)
type MyTreeNode struct {
	*tree.Node
}

func (myNode *MyTreeNode) preOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := MyTreeNode{myNode.Left}
	right := MyTreeNode{myNode.Right}

	left.preOrder()
	right.preOrder()
	myNode.Print()
}

func (myNode *MyTreeNode) Traverse() {
	fmt.Println("This is a shadowed method")
}

func main() {
	// root := tree.Node{Val: 3}
	root := MyTreeNode{&tree.Node{Val: 3}} // Embedding

	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetVal(4)

	// root.print()
	// fmt.Println()

	// root.setVal(4)
	// root.print()
	// fmt.Println()

	// var pRoot *treeNode // nil指针
	// pRoot.setVal(200)
	// // pRoot.print()	// error
	// pRoot = &root
	// pRoot.setVal(300)
	// pRoot.print()

	// root.Traverse()	// 重载函数
	root.Node.Traverse()
	fmt.Println()

	// myRoot := MyTreeNode{&root}
	// myRoot.preOrder()
	root.preOrder() // Embedding
	fmt.Println()
	// nodes := []treeNode{
	// 	{val: 3},
	// 	{},
	// 	{6, nil, nil},
	// }
	// fmt.Println(nodes)

	// // Embedding与继承的区别:无法用父类指针指向子类
	// // cannot use &root (value of type *MyTreeNode) as *tree.Node value in
	// var baseRoot *tree.Node
	// baseRoot = &root

}
