package main

import (
	"fmt"
	"github.com/lucming/practice/tree"
)

func main() {
	fmt.Println("hello world")
	p1 := tree.NewNode(5)
	p2 := tree.NewNode(3)
	p3 := tree.NewNode(6)
	p4 := tree.NewNode(2)
	p5 := tree.NewNode(4)
	p1.Left = p2
	p1.Right = p3
	p2.Left = p4
	p2.Right = p5
	p4.Left = tree.NewNode(1)

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
	root := tree.Build([]int{1, 3, 2, 5, 6, 4}, []int{1, 2, 3, 4, 5, 6})
	fmt.Println("level travel:", tree.Bfs(root))
	fmt.Println("之字打印：", tree.LevelOrder(root))
	fmt.Println(tree.BinaryTreePaths1(p1))
	fmt.Println("is same:", tree.IsSame1(tree.NewNode(1), p1))
	fmt.Println("has path:", tree.HasPath(p1, 100))
	fmt.Println("find most:", tree.GetMost(p1))
	fmt.Println("find common ancestor:", tree.FindCommonAncestor(p1, p4, p5).Val)
	fmt.Println("层次遍历：", tree.Bfs(p1))
	fmt.Println("insert value:", tree.InsertNode1(p1, 7))
	fmt.Println("层次遍历：", tree.Bfs(p1))
	p := tree.ArrayToTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(tree.Bfs(p))
}
