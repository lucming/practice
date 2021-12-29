package tree

import "fmt"

type node struct {
	Val   int
	Left  *node
	Right *node
}

func NewNode(val int) *node {
	return &node{
		Val: val,
	}
}

func (n *node) SetLeft(p *node) {
	n.Left = p
}

//递归先序遍历
func (n *node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Println(n.Val)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

//递归中序遍历
func (n *node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Println(n.Val)
	n.Right.InOrder()
}

//递归后序遍历
func (n *node) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Println(n.Val)
}

//层次遍历
//思路：每出一个根节点，入当前节点的左孩子和右孩子
func Bfs(root *node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*node{root}
	for len(queue) > 0 {
		len := len(queue)
		for len > 0 {
			len--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			result = append(result, queue[0].Val)
			queue = queue[1:]
		}
	}

	return result
}

//先序非递归
//思路：一直入左孩子，左孩子为空，开始遍历根的右孩子
func PreOrder(root *node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*node, 0)
	for len(queue) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1].Right
		queue = queue[:len(queue)-1]
	}

	return result
}

//中序非递归
//思路：借助队列，一直入左孩子，当根结点为空（走到叶子节点），队列中最后一个节点就是要出的节点，然后开始遍历右孩子
func Inorder(root *node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*node, 0)
	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		result = append(result, queue[len(queue)-1].Val)
		root = queue[len(queue)-1].Right
		queue = queue[:len(queue)-1]
	}

	return result
}

//后续非递归
func PostOrder(root *node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*node, 0)
	var lastVisted *node
	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		n := queue[len(queue)-1]
		if n.Right == nil || n.Right == lastVisted {
			result = append(result, n.Val)
			queue = queue[:len(queue)-1]
			lastVisted = n
		} else {
			root = n.Right
		}
	}
	return result
}

//镜像树
//思路：前序遍历的时候交换左右孩子的数据节点
func ReverseBinary(root *node) *node {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	ReverseBinary(root.Left)
	ReverseBinary(root.Right)

	return root
}

//非递归
func ReverseBinary1(root *node) *node {
	if root == nil {
		return nil
	}

	stack := make([]*node, 0)
	p := root
	for len(stack) > 0 || p != nil {
		for p != nil {
			p.Left, p.Right = p.Right, p.Left
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		p = p.Right
	}

	return root
}
