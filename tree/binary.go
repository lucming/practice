package tree

import "fmt"

type node struct {
	Val int
	Left *node
	Right *node
}

func NewNode(val int)*node{
	return &node{
		Val:   val,
	}
}

func (n *node)SetLeft(p *node){
	n.Left=p
}

//递归先序遍历
func (n *node)PreOrder(){
	if n==nil{
		return
	}
	fmt.Println(n.Val)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

//递归中序遍历
func (n *node)InOrder(){
	if n==nil{
		return
	}
	n.Left.InOrder()
	fmt.Println(n.Val)
	n.Right.InOrder()
}

//递归后序遍历
func (n *node)PostOrder(){
	if n==nil{
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Println(n.Val)
}





