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

//判断是否为镜像数（树的左右孩子是否堆成）
func IsMirror(root *node) bool {
	return mirror(root.Left, root.Right)
}

func mirror(left *node, right *node) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return mirror(left.Left, right.Right) && mirror(right.Left, left.Right)
}

func IsMirror1(root *node) bool {
	queue := make([]*node, 0)
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		queue = append(queue, left.Left, right.Right, right.Left, left.Right)
	}

	return true
}

func Depth(root *node) int {
	if root == nil {
		return 0
	}

	leftDepth := Depth(root.Left)
	RightDepth := Depth(root.Right)
	maxDepth := leftDepth
	if RightDepth > leftDepth {
		maxDepth = RightDepth
	}

	return maxDepth + 1
}

func Depth1(root *node) int {
	level := 0
	queue := make([]*node, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			p := queue[0]
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
			queue = queue[1:]
		}
		level++
		l = len(queue)
	}

	return level
}

//二叉树节点数
func CountNode(root *node) int {
	if root == nil {
		return 0
	}

	result := 1
	if root.Left != nil {
		result += CountNode(root.Left)
	}
	if root.Right != nil {
		result += CountNode(root.Right)
	}

	return result
}

//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	leftH, rightH := 0, 0
//	leftNode := root.Left
//	rightNode := root.Right
//	for leftNode != nil {
//		leftNode = leftNode.Left
//		leftH++
//	}
//	for rightNode != nil {
//		rightNode = rightNode.Right
//		rightH++
//	}
//	if leftH == rightH {
//		return (2 << leftH) - 1
//	}
//	return countNodes(root.Left) + countNodes(root.Right) + 1
//}

func CountNode1(root *node) int {
	if root == nil {
		return 0
	}

	leftHight, rightHight := 0, 0
	leftNode, rightNode := root.Left, root.Right

	for leftNode != nil {
		leftNode = leftNode.Left
		leftHight++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightHight++
	}
	if leftHight == rightHight {
		return 2<<leftHight - 1
	}

	return CountNode1(root.Left) + CountNode1(root.Right) + 1
}
