package tree

import (
	"fmt"
	"strconv"
)

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

//判断二叉树是否为平衡二叉树（每个子树的左右孩子高度差<=1）
func IsBalance(root *node) bool {
	if root == nil {
		return true
	}

	if !IsBalance(root.Left) || !IsBalance(root.Right) {
		return false
	}

	leftHight := treeHight(root.Left) + 1
	rightHight := treeHight(root.Right) + 1
	if abs(leftHight, rightHight) > 1 {
		return false
	}

	return true
}

//求二叉树所有路径，递归&回溯
func BinaryTreePaths(root *node) []string {
	result := make([]string, 0)
	var travel func(*node, string)
	travel = func(root *node, s string) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			result = append(result, s+strconv.Itoa(root.Val))
			return
		}
		if root.Left != nil {
			travel(root.Left, s+strconv.Itoa(root.Val)+"->")
		}
		if root.Right != nil {
			travel(root.Right, s+strconv.Itoa(root.Val)+"->")
		}
	}

	travel(root, "")
	return result
}

//非递归求二叉树路径
func BinaryTreePaths1(root *node) []string {
	result := make([]string, 0)
	paths := make([]string, 0)
	stack := make([]*node, 0)
	if root != nil {
		paths = append(paths, "")
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		lenStack := len(stack)
		cur := stack[lenStack-1]
		path := paths[lenStack-1]
		stack = stack[:lenStack-1]
		paths = paths[:lenStack-1]
		if cur.Left == nil && cur.Right == nil {
			result = append(result, path+strconv.Itoa(cur.Val))
			continue
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			paths = append(paths, path+strconv.Itoa(cur.Val)+"->")
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
			paths = append(paths, path+strconv.Itoa(cur.Val)+"->")
		}
	}

	return result
}

//计算树的高度(左右自孩子出生高度较高的+1)
func treeHight(root *node) int {
	if root == nil {
		return 0
	}

	return max(treeHight(root.Left), treeHight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//前序+中序 构建树
//思路：前序的第一个节点是根，然后拿着这个值在终须中找根节点的位置，这样就确定左右子孩子了，然后继续递归
func BuildTree(preOrder []int, inorder []int) *node {
	if len(preOrder) < 1 || len(inorder) < 1 {
		return nil
	}

	rootIndex := findNextRoot(inorder, preOrder[0])
	return &node{
		Val:   preOrder[0],
		Left:  BuildTree(preOrder[1:rootIndex+1], inorder[:rootIndex]),
		Right: BuildTree(preOrder[rootIndex+1:], inorder[rootIndex+1:]),
	}
}

func findNextRoot(inorder []int, target int) int {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}

	return -1
}

//后序+中序 构建树
func Build(postorder []int, inorder []int) *node {
	if len(postorder) < 1 || len(inorder) < 1 {
		return nil
	}
	nextRootIndex := findNextRoot(inorder, postorder[len(postorder)-1])
	return &node{
		Val:   postorder[len(postorder)-1],
		Left:  Build(postorder[:nextRootIndex], inorder[:nextRootIndex]),
		Right: Build(postorder[nextRootIndex:len(postorder)-1], inorder[nextRootIndex+1:]),
	}
}

//判断两棵树是否相等
func IsSame(root1 *node, root2 *node) bool {
	if root1 == nil && root2 != nil {
		return false
	} else if root2 == nil && root1 != nil {
		return false
	} else if root1 == nil && root2 == nil {
		return true
	} else if root1.Val != root2.Val {
		return false
	}

	return IsSame(root1.Left, root2.Left) && IsSame(root1.Right, root2.Right)
}

//非递归判断两棵树是否相等
func IsSame1(root1 *node, root2 *node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	queue := make([]*node, 0)
	queue = append(queue, root1, root2)
	for len(queue) > 0 {
		p1 := queue[0]
		p2 := queue[1]
		queue = queue[2:]
		if p1 == nil && p2 == nil {
			continue
		}
		if (p1 == nil || p2 == nil) || (p1.Val != p2.Val) {
			return false
		}
		queue = append(queue, p1.Left, p1.Right, p2.Left, p2.Right)
	}

	return true
}

// 判断树是否存在一条从根到叶子节点的路径，使得途径的所有节点的value之和等于某个值
// 思路：先序遍历，判断左右子树中是否存在这样的路径
func HasPath(root *node, sum int) bool {
	if root == nil {
		return false
	}

	return findPath(root, 0, sum)
}

func findPath(root *node, cur, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val+cur == sum {
		return true
	}

	return findPath(root.Left, cur+root.Val, sum) || findPath(root.Right, cur+root.Val, sum)
}

//获取树中出现的众数
func GetMost(root *node) []int {
	result := make([]int, 0)
	count, max := 0, 0
	var pre *node
	var travel func(*node)
	travel = func(root *node) {
		if root == nil {
			return
		}
		travel(root.Left)
		if pre != nil && pre.Val == root.Val {
			count++
		} else {
			count = 1
		}

		if count >= max {
			if count > max && len(result) > 0 {
				result = []int{root.Val}
			} else {
				result = append(result, root.Val)
			}
			max = count
		}
		pre = root
		travel(root.Right)
	}

	travel(root)
	return result
}

//寻找两节点公共祖先
//思路：后续遍历
func FindCommonAncestor(root, p, q *node) *node {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	left := FindCommonAncestor(root.Left, p, q)
	right := FindCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	if right != nil {
		return right
	}

	return nil
}
