package main

import (
	"fmt"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		fmt.Printf("%d ", node.Val)
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.TraverseFunc(f)
	f(node)
	node.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		// 将node的信息n传给channel out,类似于yeild
		node.TraverseFunc(func(n *Node) {
			out <- n
		})
		close(out)
	}()
	return out
}

func main() {
	var root Node

	// 遍历节点
	root.Traverse()

	// 统计节点数量
	nodeCount := 0
	root.TraverseFunc(func(node *Node) {
		nodeCount++
	})
	fmt.Println("Node Count: ", nodeCount)

	// 统计maxVal
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Val > maxNode {
			maxNode = node.Val
		}
	}
	fmt.Println("Max node value: ", maxNode)
}
