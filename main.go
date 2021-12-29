package main

import (
	"fmt"
	"github.com/lucming/practice/tree"
)

func main() {
	fmt.Println("hello world")
	p1 := tree.NewNode(1)
	p2 := tree.NewNode(2)
	p3 := tree.NewNode(3)
	p4 := tree.NewNode(4)
	p5 := tree.NewNode(5)
	p1.Left = p2
	p1.Right = p3
	p2.Left = p4
	p2.Right = p5

	//p1.PostOrder()
	result := tree.Bfs(p1)
	fmt.Println("层次遍历：", result)
	result = tree.PreOrder(p1)
	fmt.Println("先序遍历：", result)
	result = tree.Inorder(p1)
	fmt.Println("中序遍历：", result)
	result = tree.PostOrder(p1)
	fmt.Println("后序遍历：", result)

	tree.ReverseBinary1(p1)
	fmt.Println("层次遍历：", tree.Bfs(p1))
}
