package main

import (
	"fmt"
	"github.com/lucming/practice/tree"
)

func main() {
	fmt.Println("hello world")
	p1 := tree.NewNode(1)
	p2 := tree.NewNode(2)
	p3 := tree.NewNode(2)
	p4 := tree.NewNode(4)
	p5 := tree.NewNode(4)
	p1.Left = p2
	p1.Right = p3
	p2.Left = p4
	p2.Right = p5
	p4.Left = tree.NewNode(100)

	//p1.PostOrder()
	result := tree.Bfs(p1)
	fmt.Println("层次遍历：", result)
	result = tree.PreOrder(p1)
	fmt.Println("先序遍历：", result)
	result = tree.Inorder(p1)
	fmt.Println("中序遍历：", result)
	result = tree.PostOrder(p1)
	fmt.Println("后序遍历：", result)

	tree.ReverseBinary(p1)
	fmt.Println("层次遍历：", tree.Bfs(p1))
	fmt.Println("is mirror:", tree.IsMirror1(p1))
	fmt.Println("depth:", tree.Depth1(p1))
	fmt.Println("count node:", tree.CountNode1(p1))
	fmt.Println("is balance:", tree.IsBalance(p1))
	//root := tree.BuildTree([]int{1, 2, 4, 100, 4, 2}, []int{100, 4, 2, 4, 1, 2})
	root := tree.Build([]int{100, 4, 4, 2, 2, 1}, []int{100, 4, 2, 4, 1, 2})
	fmt.Println("level travel:", tree.Bfs(root))
}
