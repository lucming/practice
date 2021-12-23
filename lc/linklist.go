package lc

import "fmt"

type node struct {
	data int
	next *node
}

func (n *node) headInsert(d int) {
	s := &node{
		data: d,
	}
	if n == nil {
		n = s
	} else {
		s.next = n.next
		n.next = s
	}
}

func (n *node) tailInsert(d int) {
	p := n
	for ; p.next != nil; p = p.next {
		p = p.next
	}
	s := &node{data: d}
	s.next = p.next
	p.next = s
}

//头插法实现链表逆置
func reverse(head *node) {
	if head == nil || head.next == nil || head.next.next == nil {
		return
	}

	p := head.next
	head.next = nil
	var q *node
	for p != nil {
		q = p.next
		p.next = head.next
		head.next = p
		p = q
	}
}

//结点不动，第一个结点的数据域与最后一个结点的数据域的值交换，
//第二个结点的数据域与倒数第二个结点的数据域交换，以此类推
//时间复杂度O(n^2)
func reverse1(head *node) {
	p := head.next
	var q *node
	for q = head; q.next != nil; q = q.next {
		for i := 0; i < getLength(head)/2; i++ {
			p.data, q.data = q.data, p.data
			p = p.next
			q = getPri(head, q)
		}
	}
}

//改变指针的指向，后->前，头指针指向最后一个结点 //时间复杂的O(n)
func reverse2(head *node) {
	if head == nil || head.next == nil || head.next.next == nil {
		return
	}

	p := head.next
	q := p.next
	p.next = nil
	var s *node

	for q != nil {
		s = q.next
		q.next = p
		p = q
		q = s
	}
	head.next = p
}

func getLength(head *node) int {
	count := 0
	for head != nil {
		head = head.next
		count++
	}

	return count
}

func getPri(head *node, find *node) *node {
	if head == nil {
		return head
	}

	for head.next != nil && head.next != find {
		head = head.next
	}

	return head
}

func (n *node) show() {
	for p := n; p != nil; p = p.next {
		fmt.Println(p.data)
	}
}
