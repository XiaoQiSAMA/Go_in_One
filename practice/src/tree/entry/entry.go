package main

import "learngo/src/tree"

func main() {
	root := tree.Node{Val: 3}
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

	root.Traverse()
	// nodes := []treeNode{
	// 	{val: 3},
	// 	{},
	// 	{6, nil, nil},
	// }
	// fmt.Println(nodes)

}
